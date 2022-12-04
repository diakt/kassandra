package seed

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "my_password"
	dbname = "temp"

)

func Seed(){
	fmt.Printf("In Seed\n")
	
	conn_details := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	connection, err := sql.Open("postgres", conn_details)
	if err!=nil{
		fmt.Println("Something in connection open died.")
		panic(err)
	}
	defer connection.Close()

	err = connection.Ping()
	if err!=nil{
		panic(err)
	}
	fmt.Printf("Connected with %s\n", conn_details)

	//Actual database seeding
	// createCommand := `
	// CREATE TABLE practices (
	// 	practice_id SERIAL PRIMARY KEY,
	// 	practice_date TEXT,
	// 	workout_id TEXT,
	// 	boat1_id INT,
	// 	boat2_id INT
	//   );
	  
	//   CREATE TABLE workouts (
	// 	workout_id SERIAL PRIMARY KEY,
	// 	intensity INT,
	// 	volume INT,
	// 	structure TEXT
	//   );
	  
	//   CREATE TABLE boats (
	// 	boat_id SERIAL PRIMARY KEY,
	// 	bow_athlete_id INT,
	// 	two_athlete_id INT,
	// 	three_athlete_id INT,
	// 	four_athlete_id INT,
	// 	five_athlete_id INT,
	// 	six_athlete_id INT,
	// 	seven_athlete_id INT,
	// 	eight_athlete_id INT,
	// 	cox_athlete_id INT,
	// 	shell_id INT
	//   );
	  
	//   CREATE TABLE athletes (
	// 	ath_id SERIAL PRIMARY KEY,
	// 	ath_name CHAR,
	// 	ath_2k_watt TEXT,
	// 	ath_6k_watt INT
	//   );
	  
	//   CREATE TABLE coxswains (
	// 	cox_id SERIAL PRIMARY KEY,
	// 	cox_name CHAR,
	// 	cox_weight INT
	//   );
	  
	//   CREATE TABLE shells (
	// 	shell_id SERIAL PRIMARY KEY,
	// 	shell_name CHAR
	//   );
	  
	  
	//   INSERT INTO athletes (ath_id, ath_name, ath_2k_watt, ath_6k_watt) 
	//   VALUES 
	// 	  (DEFAULT,'B', 101, 80),
	// 	  (DEFAULT,'B', 102, 81),
	// 	  (DEFAULT,'C', 103, 82),
	// 	  (DEFAULT,'D', 104, 83),
	// 	  (DEFAULT,'E', 105, 84),
	// 	  (DEFAULT,'F', 106, 85),
	// 	  (DEFAULT,'G', 107, 86),
	// 	  (DEFAULT,'H', 108, 87),
	// 	  (DEFAULT,'I', 109, 88),
	// 	  (DEFAULT,'J', 110, 89),
	// 	  (DEFAULT,'K', 111, 90),
	// 	  (DEFAULT,'L', 112, 91),
	// 	  (DEFAULT,'M', 113, 92),
	// 	  (DEFAULT,'N', 114, 93),
	// 	  (DEFAULT,'O', 115, 94),
	// 	  (DEFAULT,'P', 116, 95);
	  
	//   INSERT INTO coxswains (cox_id, cox_name, cox_weight)
	//   VALUES 
	// 	  (DEFAULT, 'O', 115),
	// 	  (DEFAULT, 'T', 120);
	  
	//   INSERT INTO shells (shell_id, shell_name)
	//   VALUES 
	// 	  (DEFAULT,'S'),
	// 	  (DEFAULT,'D');
	  
	//   INSERT INTO boats (boat_id, bow_athlete_id, two_athlete_id,three_athlete_id,four_athlete_id,five_athlete_id,six_athlete_id,seven_athlete_id,eight_athlete_id,cox_athlete_id,shell_id)
	//   VALUES 
	// 	  (DEFAULT, 1, 3, 5, 8, 9, 11, 14, 16, 1, 1),
	// 	  (DEFAULT, 2, 4, 6, 7, 10, 12, 13, 15, 2, 2);
	// `
	// _, err = connection.Exec(createCommand)
	// if err!=nil{
	// 	fmt.Println("Insertion didn't work.")
	// 	panic(err)
	// }




	// insCommand := `INSERT INTO athletes (id, ath_name, ath_2k_watt, ath_6k_watt) VALUES (DEFAULT,'A', 101, 103);`
	
	
	
	fmt.Printf("Exiting seed\n")




}














