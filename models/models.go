package models



type practice struct{
	Practice_ID string `json:"practice_id"`
	Practice_date string `json:"practice_date"`
	Workout_ID int `json:"workout_id"`
	Boat1 int `json:"boat1_id"`
	Boat2 int `json:"boat2_id"`
}

type workout struct {
	Workout_ID int `json:"workout_id"`
	Intensity int `json:"intensity"`
	Volume int `json:"volume"`
	Structure string `json:"structure"`
}

type boat struct {
	Boat_ID int `json:"boat_id"`
	Bow int `json:"bow_athlete_id"`
	Two int `json:"two_athlete_id"`
	Three int `json:"three_athlete_id"`
	Four int `json:"four_athlete_id"`
	Five int `json:"five_athlete_id"`
	Six int `json:"six_athlete_id"`
	Seven int `json:"seven_athlete_id"`
	Eight int `json:"eight_athlete_id"`
	Cox int `json:"cox_athlete_id"`
	Shell int `json:"shell_id"`
}

type athlete struct {
	Athlete_ID int `json:"athlete_id"`
	Ath_name string `json:"ath_name"`
	Ath_2k_watt int `json:"ath_2k_watt"`
	Ath_6k_watt int `json:"ath_6k_watt"`

}