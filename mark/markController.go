package mark

import (
	"errors"
	"strconv"
	"time"
)

type CheckCategory string

const (
	GYM    CheckCategory = "Gym"
	CODING CheckCategory = "Coding"
)

type CheckMark struct {
	ID       string        `json:"id"`
	UserID   string        `json:"user_id"`
	Time     string        `json:"time"`
	Content  string        `json:"content"`
	Category CheckCategory `json:"category"`
}

type Controller interface {
	CreateMark(mark CheckMark) error
	ListMarks() ([]CheckMark, error)
	GetMark(id string) (CheckMark, error)
	DeleteMark(id string) error
}

func NewController() Controller {
	list := make([]CheckMark, 0)
	return &controller{
		checkMarkList: list,
	}
}

type controller struct {
	checkMarkList []CheckMark
}

func (c *controller) CreateMark(mark CheckMark) error {
	if mark.UserID == "" {
		return errors.New("requires valid user")
	}

	if mark.ID == "" {
		mark.ID = strconv.Itoa(len(c.checkMarkList) + 1)
	}

	if mark.Time == "" {
		mark.Time = time.Now().Format("2006-01-02 15:04:05.999999999 -0700 MST")
	}

	if mark.Category == "" {
		mark.Category = GYM
	}

	c.checkMarkList = append(c.checkMarkList, mark)

	return nil
}

func (c *controller) ListMarks() ([]CheckMark, error) {
	return c.checkMarkList, nil
}

func (c *controller) GetMark(id string) (CheckMark, error) {
	for _, mark := range c.checkMarkList {
		if mark.ID == id {
			return mark, nil
		}
	}
	return CheckMark{}, errors.New("failed to get mark with ID: " + id)
}

func (c *controller) DeleteMark(id string) error {
	for index, mark := range c.checkMarkList {
		if mark.ID == id {
			c.checkMarkList = append(c.checkMarkList[:index], c.checkMarkList[index+1:]...)
			return nil
		}
	}
	return errors.New("Failed to delete mark with ID: " + id)
}
