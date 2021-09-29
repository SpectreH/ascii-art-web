package art

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// Every character will have own id by ascii-table and own form. The last one we save in 2d array
type Banner struct {
	id          int
	asciiSymbol [8][]rune
}

func PrintAscii(message string, textures string) string {

	var bannerTemplateList []Banner = LoadTemplatePack(textures)
	var transformedInput [][]rune = TransformInput(message)
	var bannersToPrint [][]Banner = CollectNeededBanners(transformedInput, bannerTemplateList)
	SaveBannerInToFile("./output/result.txt", bannersToPrint)

	return (PrintBanners(bannersToPrint))
}

// Loads all symbols from text file with ascii-characters and returns them in array
func LoadTemplatePack(textures string) []Banner {
	var result []Banner

	texturesPath := "./textures/" + textures + ".txt"

	file, err := ioutil.ReadFile(texturesPath)
	CheckFile(err)

	text := TranslateToRuneSlice(file)
	var bannerToApply Banner
	var textIndex int

	if text[0] == 10 {
		textIndex = 1
	} else {
		textIndex = 2
	}

	for i := 32; i < 127; i++ {
		var tempArr [8][]rune

		for k := 0; k < 8; k++ {
			for l := 0; l < 32; l++ {

				if text[textIndex] == 13 {
					if text[textIndex+1] == 10 {
						textIndex = textIndex + 2
						break
					}
				} else if text[textIndex] == 10 {
					textIndex++
					break
				}

				tempArr[k] = append(tempArr[k], text[textIndex])
				textIndex++
			}
		}

		bannerToApply.id = i
		bannerToApply.asciiSymbol = tempArr
		result = append(result, bannerToApply)

		if i != 126 && text[textIndex] == 13 {
			textIndex = textIndex + 2
		} else {
			textIndex++
		}
	}

	return result
}

// Transform string text to 2d rune array. We separate chars by rows (If we have "/n" it means that we add new row to array)
func TransformInput(text string) [][]rune {
	var result [][]rune

	textInRune := []rune(text)
	currentLine := 0

	for i := 0; i < len(textInRune); i++ {
		if textInRune[i] == 13 && i+1 < len(textInRune) {
			if textInRune[i+1] == 10 {
				result = append(result, nil)
				currentLine++
				i++
				continue
			}
		}

		if result == nil {
			result = append(result, nil)
		} else if result[0] == nil && len(result) == currentLine {
			result = append(result, nil)
		}

		result[currentLine] = append(result[currentLine], textInRune[i])
	}

	return result
}

// With transformed string - here we try to find needed ascii-symbol and save it in to 2d array. Here we also seperate ascii-symbol by rows like in previous function
func CollectNeededBanners(charList [][]rune, bannerList []Banner) [][]Banner {
	var result [][]Banner

	for i := 0; i < len(charList); i++ {
		if result == nil {
			result = append(result, nil)
		}

		for k := 0; k < len(charList[i]); k++ {
			for m := 0; m < len(bannerList); m++ {
				if charList[i][k] == rune(bannerList[m].id) {
					result[i] = append(result[i], bannerList[m])
				} else {
					continue
				}
			}
		}

		if len(charList) != i+1 {
			result = append(result, nil)
		}
	}

	return result
}

// Prints all ascii-characters by our 2d banner array what we have built. Nil array means new-line
func PrintBanners(banners [][]Banner) string {
	var result string

	for i := 0; i < len(banners); i++ {
		if banners[i] == nil {
			// fmt.Fprintf(w, "\n")
			result = result + "\n"
			fmt.Println()
			continue
		}

		for k := 0; k < 8; k++ {
			for d := 0; d < len(banners[i]); d++ {
				// fmt.Fprintf(w, string(banners[i][d].asciiSymbol[k]))
				result = result + string(banners[i][d].asciiSymbol[k])
				fmt.Print(string(banners[i][d].asciiSymbol[k]))
			}
			// fmt.Fprintf(w, "\n")
			result = result + "\n"
			fmt.Println()
		}
	}

	return result
}

func SaveBannerInToFile(fileName string, bannersToSave [][]Banner) {
	_ = os.Remove(fileName)
	file, _ := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)

	dataWriter := bufio.NewWriter(file)
	for i := 0; i < len(bannersToSave); i++ {
		if bannersToSave[i] == nil {
			_, _ = dataWriter.WriteString("\n")
			continue
		}

		for k := 0; k < 8; k++ {
			for d := 0; d < len(bannersToSave[i]); d++ {
				_, _ = dataWriter.WriteString(string(bannersToSave[i][d].asciiSymbol[k]))
			}
			_, _ = dataWriter.WriteString("\n")
		}
	}

	_, _ = dataWriter.WriteString("\n")
	dataWriter.Flush()
	file.Close()
}

func CheckFile(e error) {
	if e != nil {
		fmt.Println("Error. Missing file")
		os.Exit(1)
	}
}

func TranslateToRuneSlice(bytes []byte) []rune {
	var text []rune

	for i := range bytes {
		text = append(text, rune(bytes[i]))
	}

	return text
}
