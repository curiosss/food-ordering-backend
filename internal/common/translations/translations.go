package translations

func GetText(key, locele string) string {
	ts := texts[key][locele]
	if ts != "" {
		return ts
	}

	return key
}

var texts = map[string]map[string]string{
	"userExists": {
		"tk": "Bu nomer eýýäm hasaba alnan",
		"ru": "Вы уже зарегистрированы с этим номером",
		"en": "You already registered with this number",
	},
}
