package textHandler

import (
	"errors"
	"strings"
)

const url = "https://cdn.simpleicons.org/"

func getIcon(language string) (string, error) {
	switch strings.ToLower(language) {
	case "typescript":
		return generateImgTag("typescript", "3178C6"), nil
	case "javascript":
		return generateImgTag("javascript", "F7DF1E"), nil
	case "css":
		return generateImgTag("css3", "1572B6"), nil
	case "html":
		return generateImgTag("html5", "E34F26"), nil
	case "go":
		return generateImgTag("go", "00ADD8"), nil
	case "python":
		return generateImgTag("python", "3776AB"), nil
	case "java":
		return generateImgTag("java", "007396"), nil
	case "c#":
		return generateImgTag("csharp", "239120"), nil
	case "c++":
		return generateImgTag("cplusplus", "00599C"), nil
	case "c":
		return generateImgTag("c", "A8B9CC"), nil
	case "ruby":
		return generateImgTag("ruby", "CC342D"), nil
	case "php":
		return generateImgTag("php", "777BB4"), nil
	case "markdown":
		return generateImgTag("markdown", "fff"), nil
	case "bash":
		return generateImgTag("gnubash", "fff"), nil
	case "yaml":
		return generateImgTag("yaml", "fff"), nil
	case "json":
		return generateImgTag("carrd", "fff"), nil
	case "vue":
		return generateImgTag("vuedotjs", "4FC08D"), nil
	case "text":
		return generateImgTag("academia", "fff"), nil
	default:
		return "", errors.New("not found")
	}
}

func generateImgTag(language string, color string) string {
	return url + language + "/" + color
}
