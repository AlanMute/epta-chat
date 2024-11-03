package repository

import (
	"fmt"
	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/jinzhu/gorm"
)

type ChatRepo struct {
	db *gorm.DB
}

func NewChatPostgres(db *gorm.DB) *ChatRepo {
	return &ChatRepo{
		db: db,
	}
}

func (r *ChatRepo) Add(name string, isDirect bool, ownerId uint64, members []uint64) (uint64, error) {
	var err error
	if isDirect {
		var potentialChats []core.Chat
		if err := r.db.Where("is_direct = ? AND owner_id IN (?)", true, members).Find(&potentialChats).Error; err != nil {
			return 0, err
		}

		for _, chat := range potentialChats {
			var chatMemberIds []uint64
			r.db.Model(&core.ChatMembers{}).Where("chat_id = ?", chat.ID).Pluck("member_id", &chatMemberIds)

			if equalMembers(chatMemberIds, members) {
				return chat.ID, nil
			}
		}
	}

	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	newChat := core.Chat{
		Name:     name,
		IsDirect: isDirect,
		OwnerId:  ownerId,
	}

	if err = tx.Create(&newChat).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	for _, memberId := range members {
		var existingMember core.ChatMembers
		if err = tx.Where("chat_id = ? AND member_id = ?", newChat.ID, memberId).First(&existingMember).Error; err == nil {
			continue
		}

		newMember := core.ChatMembers{
			MemberId: memberId,
			ChatId:   newChat.ID,
		}

		if err = tx.Create(&newMember).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	return newChat.ID, tx.Commit().Error
}

func (r *ChatRepo) AddMember(ownerId, chatId uint64, members []uint64) error {
	var chat core.Chat

	if result := r.db.Where("id = ?", chatId).First(&chat); result.Error != nil {
		return result.Error
	}

	if ownerId != chat.OwnerId {
		return fmt.Errorf("not an owner of this chat")
	}

	if chat.IsDirect {
		return fmt.Errorf("direct chat cannot has more than 2 members")
	}

	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, memberId := range members {
		var existingMember core.ChatMembers
		if err := tx.Where("chat_id = ? AND member_id = ?", chatId, memberId).First(&existingMember).Error; err == nil {
			continue
		}

		newMember := core.ChatMembers{
			MemberId: memberId,
			ChatId:   chatId,
		}

		if err := tx.Create(&newMember).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (r *ChatRepo) Delete(userId, chatId uint64) error {
	var chat core.Chat

	if result := r.db.Where("id = ?", chatId).First(&chat); result.Error != nil {
		return result.Error
	}

	if chat.IsDirect {
		var member core.ChatMembers

		if result := r.db.Where("chat_id = ? AND member_id = ?", chatId, userId).First(&member); result.Error != nil {
			if gorm.IsRecordNotFoundError(result.Error) {
				return fmt.Errorf("not a member of the direct chat")
			}
			return result.Error
		}
	} else if userId != chat.OwnerId {
		return fmt.Errorf("not an owner of the chat")
	}

	if result := r.db.Delete(&chat); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ChatRepo) GetById(userId, chatId uint64) (core.Chat, error) {
	var chat core.Chat

	if result := r.db.Where("id = ?", chatId).First(&chat); result.Error != nil {
		return core.Chat{}, result.Error
	}

	return r.setChatName(chat, userId)
}

func (r *ChatRepo) GetAll(userId uint64) ([]core.Chat, error) {
	var userChats []core.ChatMembers

	if result := r.db.Where("member_id = ?", userId).Find(&userChats); result.Error != nil {
		return nil, result.Error
	}

	chatIDs := make([]uint64, len(userChats))
	for i, userChat := range userChats {
		chatIDs[i] = userChat.ChatId
	}

	var chats []core.Chat
	if result := r.db.Where("id IN (?)", chatIDs).Find(&chats); result.Error != nil {
		return nil, result.Error
	}

	var err error
	for i, chat := range chats {
		chats[i], err = r.setChatName(chat, userId)
		if err != nil {
			return nil, err
		}
	}

	return chats, nil
}

func (r *ChatRepo) GetMembers(userId, chatId uint64) ([]core.UserInfo, error) {
	var member core.ChatMembers
	if result := r.db.Where("chat_id = ? AND member_id = ?", chatId, userId).First(&member); result.Error != nil {
		if gorm.IsRecordNotFoundError(result.Error) {
			return nil, fmt.Errorf("user is not a member of the chat")
		}
		return nil, result.Error
	}

	var chatMembers []core.ChatMembers
	if result := r.db.Where("chat_id = ?", chatId).Find(&chatMembers); result.Error != nil {
		return nil, result.Error
	}

	var userIDs []uint64
	for _, member := range chatMembers {
		userIDs = append(userIDs, member.MemberId)
	}

	var users []core.User
	if result := r.db.Where("id IN (?)", userIDs).Find(&users); result.Error != nil {
		return nil, result.Error
	}

	var usersInfo []core.UserInfo
	for _, user := range users {
		usersInfo = append(usersInfo, core.UserInfo{
			ID:       user.ID,
			Login:    user.Login,
			UserName: user.UserName,
		})
	}

	return usersInfo, nil
}

func (r *ChatRepo) setChatName(chat core.Chat, userId uint64) (core.Chat, error) {
	if chat.IsDirect {
		var otherMember core.ChatMembers

		if result := r.db.Where("chat_id = ? AND member_id != ?", chat.ID, userId).First(&otherMember); result.Error != nil {
			if gorm.IsRecordNotFoundError(result.Error) {
				return core.Chat{}, fmt.Errorf("not a member of the chat")
			}
			return core.Chat{}, result.Error
		}

		var user core.User
		if result := r.db.Where("id = ?", otherMember.MemberId).First(&user); result.Error != nil {
			return core.Chat{}, result.Error
		}

		chat.Name = user.UserName
	} else {
		var member core.ChatMembers

		if result := r.db.Where("chat_id = ? AND member_id = ?", chat.ID, userId).First(&member); result.Error != nil {
			if gorm.IsRecordNotFoundError(result.Error) {
				return core.Chat{}, fmt.Errorf("not a member of the chat")
			}
			return core.Chat{}, result.Error
		}
	}

	return chat, nil
}

func (r *ChatRepo) EnsureCommonChatExists() error {
	var commonChat core.Chat
	result := r.db.First(&commonChat, "id = ?", 1)

	if result.Error == gorm.ErrRecordNotFound {
		commonChat = core.Chat{
			Name:     "Common chat",
			IsDirect: false,
		}
		if err := r.db.Create(&commonChat).Error; err != nil {
			return err
		}
	} else {
		return result.Error
	}

	return nil
}

func equalMembers(existingMembers, newMembers []uint64) bool {
	if len(existingMembers) != len(newMembers) {
		return false
	}

	memberMap := make(map[uint64]bool)
	for _, member := range existingMembers {
		memberMap[member] = true
	}

	for _, member := range newMembers {
		if !memberMap[member] {
			return false
		}
	}

	return true
}

func (r *ChatRepo) FetchAllChatIDs() ([]uint64, error) {
	var chats []core.Chat

	if result := r.db.Find(&chats); result.Error != nil {
		return nil, result.Error
	}

	chatIDs := make([]uint64, len(chats))
	for i, chat := range chats {
		chatIDs[i] = chat.ID
	}

	return chatIDs, nil
}
