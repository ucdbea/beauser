package controllers

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"github.com/revel/revel"
	"google.golang.org/api/option"
)

type Calendar struct {
	*revel.Controller
}

func (c Calendar) Calendar() revel.Result {
	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: "https://beawebsite-86b5d.firebaseio.com/",
	}
	sa := option.WithCredentialsFile("./RealtimeDatabase_SA.json")
	app, err := firebase.NewApp(ctx, conf, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}

	//Get a database reference to our blog.
	ref := client.NewRef("MyDatabase")
	type Event struct {
		Date        string `json:"date_of_birth,omitempty"`
		description string `json:"full_name,omitempty"`
		name        string `json:"nickname,omitempty"`
	}

	usersRef := ref.Child("MyDatabase")

	err = usersRef.Set(ctx, map[string]*Event{
		"alanisawesome": {
			Date: "June 23, 1912",
			name: "Alan Turing",
		},
		"gracehop": {
			Date: "December 9, 1906",
			name: "Grace Hopper",
		},
	})
	if err != nil {
		log.Fatalln("Error setting value:", err)
	}
	return c.Render()
}

func (c Calendar) addEvent(date, name, description string) revel.Result {
	ctx := context.Background()
	sa := option.WithCredentialsFile("./RealtimeDatabase_SA.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}

	// Get a database reference to our blog.
	ref := client.NewRef("MyDatabase")
	type Event struct {
		Date        string `json:"date_of_birth,omitempty"`
		description string `json:"full_name,omitempty"`
		name        string `json:"nickname,omitempty"`
	}

	usersRef := ref.Child("users")
	err = usersRef.Set(ctx, map[string]*Event{
		"alanisawesome": {
			Date: "June 23, 1912",
			name: "Alan Turing",
		},
		"gracehop": {
			Date: "December 9, 1906",
			name: "Grace Hopper",
		},
	})
	if err != nil {
		log.Fatalln("Error setting value:", err)
	}
	return c.Redirect(Calendar.Calendar)
}

func (c Calendar) getEvent() revel.Result {
	ctx := context.Background()
	sa := option.WithCredentialsFile("./RealtimeDatabase_SA.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}

	// Get a database reference to our blog.
	ref := client.NewRef("MyDatabase")
	type Event struct {
		Date        string `json:"date_of_birth,omitempty"`
		description string `json:"full_name,omitempty"`
		name        string `json:"nickname,omitempty"`
	}

	usersRef := ref.Child("users")
	err = usersRef.Set(ctx, map[string]*Event{
		"alanisawesome": {
			Date: "June 23, 1912",
			name: "Alan Turing",
		},
		"gracehop": {
			Date: "December 9, 1906",
			name: "Grace Hopper",
		},
	})
	if err != nil {
		log.Fatalln("Error setting value:", err)
	}
	return c.Redirect(Calendar.Calendar)
}
