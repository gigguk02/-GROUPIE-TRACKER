package pkg

import (
	"artists/models"
	"net/http"
	"strconv"
	"strings"
	"unicode"
)

func Search(clientWords string, Artists []models.Artist) ([]models.Artist, int) {
	var result []models.Artist
	if _, err := strconv.Atoi(clientWords); err == nil {
		for _, inf := range Artists {
			creationDateString := strconv.Itoa(int(inf.CreationDate))
			if isContain(inf.Name, clientWords) {
				result = append(result, inf)
			} else if isContain(creationDateString, clientWords) {
				result = append(result, inf)

			}
		}
		return result, http.StatusOK

	} else if IsData(clientWords) {
		for _, inf := range Artists {
			if isContain(inf.FirstAlbum, clientWords) {
				result = append(result, inf)

			}
		}
		return result, http.StatusOK

	} else if LettersOnly(clientWords) {
		flag := true
		for index, inf := range Artists {
			if isContain(inf.Name, clientWords) {
				result = append(result, inf)
			} else {
				for i := range inf.Members {
					if flag && isContain(inf.Members[i], clientWords) {
						result = append(result, inf)
						flag = false
						break

					}
				}
				if flag {
					for key := range inf.Relation {
						if isContain(key, clientWords) {
							result = append(result, Artists[index])
							break
						}

					}

				}

			}
		}
		return result, http.StatusOK

	} else {
		for _, v := range Artists {
			if isContain(v.Name, clientWords) {
				result = append(result, v)
			}
		}
		return result, http.StatusOK

	}

}
func IsData(clientWords string) bool {
	var isData []bool
	data := "0123456789-"
	for i := range clientWords {
		for j := range data {
			if clientWords[i] == data[j] {
				isData = append(isData, true)
				break
			}
		}
	}
	if len(isData) == len(clientWords) {
		return true
	}
	return false

}
func LettersOnly(clientWords string) bool {
	var notNumber []bool
	for i := range clientWords {
		if !unicode.IsDigit(rune(clientWords[i])) {
			notNumber = append(notNumber, true)
		}
	}
	if len(notNumber) == len(clientWords) {
		return true
	}
	return false
}
func isContain(str string, client string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(client))

}
