package api

type NewUserRequest struct {
	Name          string `json:"name"`
	Age           int    `json:"age"`
	WeightGoal    string `json:"weight_goal"`
	Height        int    `json:"height"`
	Sex           string `json:"sex"`
	ActivityLevel int    `json:"activity_level"`
	Email         string `json:"email"`
}
