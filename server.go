package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Username string `json:"username"`
	DisplayName string `json:"display_name"`
	URL string `json:"url"`
	ID int32 `json:"id"`
	AvatarURL string `json:"avatar_url"`
	PremiumUser bool `json:"premium_user"`
	RegisterDate string `json:"register_date"`
	Followers int `json:"followers"`
	Following int `json:"following"`
	Bio string `json:"bio"`
	Location string `json:"location"`
	ProfileViews int `json:"profile_views"`
	Commissions int `json:"commissions"`
	Rating float32 `json:"rating"`
}

func getUsers(ctx *fiber.Ctx) error {
	// afn := User {
	// 	"afn",
	// 	"affan",
	// 	"https://venic.app/afn",
	// 	939399939,
	// 	"https://venic.app/afn.png",
	// 	false,
	// 	"2019-02-17T20:25:31Z",
	// 	10000,
	// 	0,
	// 	"hi venic",
	// 	"canada",
	// 	89982389,
	// 	3,
	// 	4.88776,
	// }

	fmt.Fprintf(ctx, "hii %s\n", ctx.Params("user"))
    return nil
}

// i dont know what im doing help
func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello Venic!")
    })

    app.Get("/users/:user", getUsers)
    app.Listen(":9989")
}