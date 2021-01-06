package controllers

import (
	"context"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/revel/revel"
	"google.golang.org/api/option"
)

type Application struct {
	*revel.Controller
}

// type FirestoreEvent struct {
// 	OldValue   FirestoreValue `json:"oldValue"`
// 	Value      FirestoreValue `json:"value"`
// 	UpdateMask struct {
// 		FieldPaths []string `json:"fieldPaths"`
// 	} `json:"updateMask"`
// }

// type FirestoreValue struct {
// 	CreateTime time.Time `json:"createTime"`
// 	Name       string    `json:"name"`
// 	UpdateTime time.Time `json:"updateTime"`
// 	Fields     User      `json:"fields"`
// }

// // This is our self-defined fields.
// // FirestoreEvent.Value.Fields = User
// type User struct {
// 	Username   StringValue  `json:"userId"`
// 	Email      StringValue  `json:"email"`
// 	DateEdited IntegerValue `json:"date_edited"`
// }

// type IntegerValue struct {
// 	IntegerValue string `json:"integerValue"`
// }

// type StringValue struct {
// 	StringValue string `json:"stringValue"`
// }

// // Simple init to have a firestore client available

// // Handles the rollback to a previous document
// func handleRollback(ctx context.Context, e FirestoreEvent) error {
// 	return errors.New("Should have rolled back to a previous version")
// }

// // The function that runs with the cloud function itself
// func HandleUserChange(ctx context.Context, e FirestoreEvent) error {
// 	// This is the data that's in the database itself
// 	newFields := e.Value.Fields
// 	oldFields := e.OldValue.Fields

// 	// As our goal is simply to check if the username has changed
// 	if newFields.Username.StringValue == oldFields.Username.StringValue {
// 		log.Printf("Bad username: %s - %s", newFields.Username.StringValue, oldFields.Username.StringValue)
// 		return handleRollback(ctx, e)
// 	}

// 	// Check if the email is the same as previously
// 	if newFields.Email.StringValue != oldFields.Email.StringValue {
// 		log.Printf("Bad email: %s - %s", newFields.Email.StringValue, oldFields.Email.StringValue)
// 		return handleRollback(ctx, e)
// 	}

// 	return nil
// }
type MyHtml string

func (r MyHtml) Apply(req *revel.Request, resp *revel.Response) {
	resp.WriteHeader(http.StatusOK, "text/html")
	resp.GetWriter().Write([]byte(r))
}
func (c Application) Index() revel.Result {
	return c.Render()
}
func (c Application) Login(email, password string, remember bool) revel.Result {
	ctx := context.Background()

	sa := option.WithCredentialsFile("./Firestore_SA.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
		"email":    email,
		"password": password,
		"remember": remember,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v ", err)
	}

	return c.Redirect(Application.Homepage)

}

func (c Application) Homepage() revel.Result {
	return c.Render()
}

func (c Application) Signin() revel.Result {
	return c.Render()
}

func (c Application) Signup() revel.Result {
	return c.Render()
}
