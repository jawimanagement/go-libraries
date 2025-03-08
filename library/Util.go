package library

import (
	"encoding/base64"
	"errors"
	"io"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	_ "image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"strconv"

	qrcode "github.com/skip2/go-qrcode"
)

var (
	ErrBucket       = errors.New("Invalid bucket!")
	ErrSize         = errors.New("Invalid size!")
	ErrInvalidImage = errors.New("Invalid image!")
)
var (
	ErrFileBucket  = errors.New("Invalid bucket!")
	ErrFileSize    = errors.New("Invalid size!")
	ErrInvalidFile = errors.New("Invalid File!")
)

func SaveImageToDisk(fileNameBase, data string) (string, error) {
	splitter := strings.Split(data, ";base64,")
	idx := strings.Index(data, ";base64,")
	if idx < 0 {
		return "", ErrInvalidImage
	}
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data[idx+8:]))
	buff := bytes.Buffer{}
	_, err := buff.ReadFrom(reader)
	if err != nil {
		return "", err
	}
	if len(splitter) > 0 && splitter[0] == "data:image/svg+xml" {
		decode := func(io.Reader) (image.Image, error) {
			return image.NewRGBA(image.Rect(0, 0, 1, 1)), nil
		}
		decodeConfig := func(io.Reader) (image.Config, error) {
			return image.Config{ColorModel: color.RGBAModel, Width: 1, Height: 1}, nil
		}
		image.RegisterFormat("svg", "<?xml ", decode, decodeConfig)
		image.RegisterFormat("svg", "<svg", decode, decodeConfig)
	}

	imgCfg, fm, err := image.DecodeConfig(bytes.NewReader(buff.Bytes()))
	if err != nil {
		return "", err
	}

	if imgCfg.Width != 750 || imgCfg.Height != 685 {
		// return "", ErrSize
	}

	fileName := fileNameBase + "." + fm
	// fmt.Println("===============================")
	// fmt.Println(fileName)
	// fmt.Println("===============================")
	err = ioutil.WriteFile(fileName, buff.Bytes(), 0644)
	if err != nil {
		log.Panic(err)
	}
	return fileName, err
}

func SaveFileToDisk(fileNameBase string, data string) (string, error) {
	idx := strings.Index(data, ";base64,")
	if idx < 0 {
		return "", ErrInvalidFile
	}
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data[idx+8:]))
	buff := bytes.Buffer{}
	_, err := buff.ReadFrom(reader)
	if err != nil {
		return "", err
	}

	fileName := fileNameBase
	ioutil.WriteFile(fileName, buff.Bytes(), 0644)

	return fileName, err
}
func IsJSON(s string) bool {
	var js interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

func FileExists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func GenUID(sequence int, lengthNumber int, charMaxLenght int) string {
	stringSequence := fmt.Sprintf("%v", sequence)
	if len(stringSequence) > lengthNumber {
		prefix := stringSequence[0 : len(stringSequence)-lengthNumber]
		suffix := stringSequence[len(stringSequence)-lengthNumber : len(stringSequence)]
		numIn, _ := strconv.Atoi(prefix)
		output := GenUIDToAlpha(numIn, "", charMaxLenght)
		if len(output) < charMaxLenght {
			iOutput := charMaxLenght - len(output)
			n := 0
			for n < iOutput {
				output = "A" + output
				n++
			}
		}
		return output + suffix
	} else {

		if len(stringSequence) < lengthNumber {
			i := 0
			ln := len(stringSequence)
			for i < lengthNumber-ln {
				stringSequence = "0" + stringSequence
				i++
			}
		}

		n := 0
		for n < charMaxLenght {
			stringSequence = "A" + stringSequence
			n++
		}
		return stringSequence
	}
	return ""
}
func GenUIDToAlpha(num int, addString string, ln int) string {
	if num < 26 {
		return string('A'+num) + addString
	} else {
		if len(addString) >= ln {
			return addString + fmt.Sprintf("%v", num)
		}
		total := num - 26
		return GenUIDToAlpha(total, string('A'-1+26)+addString, ln)
	}
}

// json encode
func JsonEncode(data interface{}) (string, error) {
	jsons, err := json.Marshal(data)
	return string(jsons), err
}

// json decode
func JsonDecode(data string) (map[string]interface{}, error) {
	var dat map[string]interface{}
	json.Unmarshal([]byte(data), &dat)
	return dat, nil
}

// parse json requst Body
func JsonReqBody(c *gin.Context) (map[string]interface{}, error) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	var dat map[string]interface{}
	json.Unmarshal([]byte(value), &dat)
	return dat, nil
}

func IsDateValue(stringDate string) bool {
	_, err := time.Parse("2006-01-02", stringDate)
	return err == nil
}
func IsDateTimeValue(stringDate string) bool {
	_, err := time.Parse("2006-01-02 15:04:05", stringDate)
	return err == nil
}
func IsTimeValue(stringDate string) bool {
	_, err := time.Parse("15:04:05", stringDate)
	return err == nil
}
func NumberFormat(number float64, decimals uint, decPoint, thousandsSep string) string {
	neg := false
	if number < 0 {
		number = -number
		neg = true
	}
	dec := int(decimals)
	// Will round off
	str := fmt.Sprintf("%."+strconv.Itoa(dec)+"F", number)
	prefix, suffix := "", ""
	if dec > 0 {
		prefix = str[:len(str)-(dec+1)]
		suffix = str[len(str)-dec:]
	} else {
		prefix = str
	}
	sep := []byte(thousandsSep)
	n, l1, l2 := 0, len(prefix), len(sep)
	// thousands sep num
	c := (l1 - 1) / 3
	tmp := make([]byte, l2*c+l1)
	pos := len(tmp) - 1
	for i := l1 - 1; i >= 0; i, n, pos = i-1, n+1, pos-1 {
		if l2 > 0 && n > 0 && n%3 == 0 {
			for j := range sep {
				tmp[pos] = sep[l2-j-1]
				pos--
			}
		}
		tmp[pos] = prefix[i]
	}
	s := string(tmp)
	if dec > 0 {
		s += decPoint + suffix
	}
	if neg {
		s = "-" + s
	}

	return s
}

func RandStringFromDb(n int, db *gorm.DB, table string, cols string) string {
	rand.Seed(time.Now().UnixNano())
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	)
	b := make([]byte, n)
	for i := 0; i < n; {
		if idx := int(rand.Int63() & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i++
		}
	}
	result := map[string]interface{}{}
	rs := db.Table(table).Select(cols).Where(cols+"=?", string(b)).Scan(result)
	if rs.RowsAffected == 0 {
		return string(b)
	}
	return RandStringFromDb(n, db, table, cols)
}

func RandString(n int) string {

	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	)
	b := make([]byte, n)
	for i := 0; i < n; {
		if idx := int(rand.Int63() & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i++
		}
	}
	return string(b)
}

func GenerateQRCode(data string, savepath string, size int) error {
	err := qrcode.WriteFile(data, qrcode.Medium, size, savepath)
	return err
}

func InArray(val interface{}, array interface{}) (index int) {
	values := reflect.ValueOf(array)

	if reflect.TypeOf(array).Kind() == reflect.Slice || values.Len() > 0 {
		for i := 0; i < values.Len(); i++ {
			if reflect.DeepEqual(val, values.Index(i).Interface()) {
				return i
			}
		}
	}

	return -1
}

func StructToJson(b interface{}) map[string]string {
	val := reflect.ValueOf(b)
	output := make(map[string]string, 0)
	for i := 0; i < val.Type().NumField(); i++ {
		// skips fields without json tag
		if tag, ok := val.Type().Field(i).Tag.Lookup("json"); ok {
			output[val.Type().Field(i).Name] = tag
		}
	}
	return output
}

func verifyPassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func dateToUnix(date string) (error, int64) {
	var res int64
	layout := "2006-01-02 15:04:05"
	if IsDateValue(date) {
		d, err := time.Parse(layout, date)
		if err != nil {
			return err, 0
		}
		res = d.Unix()
	} else if IsDateTimeValue(date) {
		dt, err := time.Parse(layout, date)
		if err != nil {
			return err, 0
		}
		res = dt.Unix()
	} else {
		return fmt.Errorf("Invalid date format"), res
	}
	return nil, res
}
