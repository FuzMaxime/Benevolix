func GenerateToken(secret, email string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "email": email,
        "exp":   time.Now().Add(time.Hour * 2).Unix(),
    })
    return token.SignedString([]byte(secret))
}

func ParseToken(secret, tokenString string) (string, error) {
    tokenString = strings.TrimPrefix(tokenString, "Bearer ")
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(secret), nil
    })

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims["email"].(string), nil
    }
    return "", err
}