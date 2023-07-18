package middleware

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"template2/internal/domain/entities"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMIddleware(c *fiber.Ctx) error {
	if strings.HasPrefix(c.Path(), "/public") {
		c.Next()
	} else {
		// c.Status(fiber.StatusUnauthorized).Send([]byte("Not authorized"))
		tokenString := c.Get("Authorization")
		splitToken := strings.Split(tokenString, "Bearer ")
		if len(splitToken) > 1 {
			tokenString = splitToken[1]
		}
		fmt.Println(tokenString)
		userid, err := ParseToken(tokenString)
		if err != nil {
			fmt.Println(err)
			c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status": false,
				"error":  err,
			})
		}
		// c.Append("user_id",userid)

		// reqBodyBytes := new(bytes.Buffer)
		// json.NewEncoder(reqBodyBytes).Encode(entities.User{Id: userid, Name: "unbeliavable"})
		fmt.Println(userid)

		c.Request().Header.Add("userId", strconv.FormatUint(uint64(userid), 10))
		c.Next()
	}
	// fmt.Println(c.GetReqHeaders())
	// c.SendStatus(fiber.StatusUnauthorized)
	return nil
}

type UserClaims struct {
	UserId uint
	jwt.RegisteredClaims
}

func ParseToken(tokenString string) (uint, error) {

	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("greencheck"), nil
	})

	if token != nil && token.Valid {
		if claims, ok := token.Claims.(*UserClaims); ok {
			return claims.UserId, nil
		}
	} else if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return 0, errors.New("That's not even a token")
		} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			return 0, errors.New("Invalid signature")
		} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			return 0, errors.New("Token is either expired or not active yet")
		}
	}
	return 0, errors.New("Couldn't handle this token:")
}

func GenerateToken(user *entities.User) string {
	tokenString := ""
	claims := UserClaims{
		user.Id,
		jwt.RegisteredClaims{
			Issuer:    "fgb",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 100)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// fmt.Println("generating token:")
	// fmt.Println(token)
	tokenString, err := token.SignedString([]byte("greencheck"))
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(tokenString)
	return tokenString
}

// func fauth(c *fiber.Ctx) error {
// 	return nil
// }
