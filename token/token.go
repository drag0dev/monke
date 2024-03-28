package token

const (
    ILLEGAL = "ILLEGAL"
    EOF     = "EOF"

    IDENT = "IDENT" // identifier
    INT   = "INT" // integer

    ASSIGN   = "="
    PLUS     = "+"
    MINUS    = "-"
    BANG     = "!"
    ASTERISK = "*"
    SLASH    = "/"

    LT = "<"
    GT = ">"

    EQ     = "=="
    NOT_EQ = "!="

    COMMA     = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    FUNCTION = "FUNCTION"
    LET      = "LET"
    TRUE     = "TRUE"
    FALSE    = "FALSE"
    IF       = "IF"
    ELSE     = "ELSE"
    RETURN   = "RETURN"

)

type TokenType string

type Token struct {
    Type TokenType
    Literal string
}
