package pkg

import (
	"artists/models"
	"net/http"
	"strconv"
	"strings"
)

func Filter(createDateStart, createDateEnd, firstAlbumStart, firstAlbumEnd, location string, numberOfMembers []string, Artists []models.Artist) ([]models.Artist, int) {
	var result []models.Artist
	isNumber := []string{createDateStart, createDateEnd, firstAlbumStart, firstAlbumEnd}
	for i := range isNumber {
		if _, err := strconv.Atoi(isNumber[i]); err != nil {
			return []models.Artist{}, http.StatusBadRequest
		}
	}
	createDateStartInt, _ := strconv.Atoi(createDateStart)
	createDateEndInt, _ := strconv.Atoi(createDateEnd)
	firstAlbumStartInt, _ := strconv.Atoi(firstAlbumStart)
	firstAlbumEndInt, _ := strconv.Atoi(firstAlbumEnd)
	numberOfMembersInt := []int{}
	for i := range numberOfMembers {
		number, err := strconv.Atoi(numberOfMembers[i])
		if err != nil {
			return []models.Artist{}, http.StatusBadRequest
		}
		numberOfMembersInt = append(numberOfMembersInt, number)

	}
	flag := false
	for i := range Artists {
		AlbumInt, err := strconv.Atoi(Artists[i].FirstAlbum[6:])
		if err != nil {
			return []models.Artist{}, http.StatusInternalServerError
		}
		if firstAlbumStartInt <= AlbumInt && AlbumInt <= firstAlbumEndInt {
			if createDateStartInt <= int(Artists[i].CreationDate) && int(Artists[i].CreationDate) <= createDateEndInt {
				if len(location) == 0 {
					flag = true

				} else {
					flag = false
					for key := range Artists[i].Relation {
						if IsContain(key, location) {
							flag = true
							break

						}
					}
				}

				if len(numberOfMembers) == 0 && flag {
					result = append(result, Artists[i])
				} else if flag {
					count := 0
					for _ = range Artists[i].Members {
						count++
					}
					for j := range numberOfMembersInt {
						if numberOfMembersInt[j] == count {
							result = append(result, Artists[i])

						}
					}

				}

			}
		}

	}

	return result, http.StatusOK

}
func IsContain(str string, client string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(client))

}
func MaxMin(Artists []models.Artist) (int, int, int, int, int) {
	maxCreateDate := int(Artists[0].CreationDate)
	minCreateDate := int(Artists[0].CreationDate)
	yearFisrtAlbum := []int{}

	for i := range Artists {
		year, err := strconv.Atoi(Artists[i].FirstAlbum[6:])
		if err != nil {
			return http.StatusBadRequest, 0, 0, 0, 0
		}
		yearFisrtAlbum = append(yearFisrtAlbum, year)
		if int(Artists[i].CreationDate) > maxCreateDate {
			maxCreateDate = int(Artists[i].CreationDate)
		}
		if int(Artists[i].CreationDate) < minCreateDate {
			minCreateDate = int(Artists[i].CreationDate)
		}
	}
	maxFirstAlbum := yearFisrtAlbum[0]
	minFirstAlbum := yearFisrtAlbum[0]
	for i := range yearFisrtAlbum {
		if yearFisrtAlbum[i] > maxFirstAlbum {
			maxFirstAlbum = yearFisrtAlbum[i]
		}
		if yearFisrtAlbum[i] < minFirstAlbum {
			minFirstAlbum = yearFisrtAlbum[i]
		}

	}
	return http.StatusOK, maxCreateDate, minCreateDate, maxFirstAlbum, minFirstAlbum
}
