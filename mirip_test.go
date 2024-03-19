package mirip_test

import (
	"testing"

	"github.com/sebarcode/mirip"
	"github.com/sebarcode/mirip/adapter"
	"github.com/smartystreets/goconvey/convey"
)

var (
	names = []string{"Adi Wijaya", "Budi Santoso", "Candra Putri", "Dewi Lestari", "Eka Sari", "Fitriani Utami", "Gita Permata", "Hadi Nugroho", "Indra Kusuma", "Joko Prabowo", "Kartika Dewi", "Lina Fitri", "Mira Puspita", "Nanda Maulana", "Oktavia Sari", "Putra Ramadhan", "Rina Novianti", "Surya Wijaya", "Tika Nurhayati", "Umar Setiawan", "Vina Cahyani", "Wulan Kusuma", "Yudi Santoso", "Zara Putri", "Aditya Purnama", "Bunga Anggraini", "Cahya Pratama", "Dian Septiani", "Erlangga Perdana", "Fita Sari", "Gilang Saputra", "Hesti Puspitasari", "Irfan Maulana", "Juwita Rahayu", "Krisna Pratama", "Lusi Fitriani", "Maman Suhendar", "Nia Sari", "Oktavian Saputra", "Putri Kartika", "Rendy Kurniawan", "Siti Rahayu", "Taufik Hidayat", "Umi Kurniati", "Vino Pratama", "Widya Putri", "Yuli Astuti", "Zaki Rahman", "Andi Saputra", "Bella Sari", "Candra Wijaya", "Dinda Fitriani", "Erwin Susanto", "Fara Putri", "Gilang Satria", "Hesti Susanti", "Indra Prasetya", "Jihan Maulidya", "Krisna Putra", "Lina Cahyani", "Mira Anggraeni", "Nanda Pratama", "Oktaviani Sari", "Putra Nugraha", "Rani Fitria", "Surya Pratama", "Tasya Putri", "Umar Kusuma", "Vina Puspita", "Winda Septiani", "Yudi Purnama", "Zara Septiani", "Agus Santoso", "Bella Cahyani", "Citra Dewi", "Dika Pratama", "Erlinda Susanti", "Fauzi Maulana", "Gita Anggraeni", "Hadi Susanto", "Intan Putri", "Joko Santoso", "Kartika Sari", "Lina Dewi", "Maulana Pratama", "Nia Fitriani", "Oka Putra", "Putri Novianti", "Rendy Setiawan", "Siti Rahmawati", "Taufik Setiawan", "Umi Septiani", "Vino Pratama", "Wulan Anggraeni", "Yani Fitriani", "Zaki Pratama", "Andi Satria", "Bella Cahaya", "Citra Sari", "Dian Putri"}

	testNameFix     = "Andi Satria"
	testNameSimilar = "Zakie D Rachman"
	testNameReverse = "Widjaya Adi"

	minimalScore = float64(0.6)
)

func TestFindLevenExact(t *testing.T) {
	convey.Convey("prepare", t, func() {
		s, ok := mirip.Compare(adapter.NewLevenshtein(), testNameFix, 1, false, names...)
		convey.So(ok, convey.ShouldBeNil)
		convey.Printf(" %s = %s\n", testNameFix, s)
	})
}

func TestFindLevenSimilar(t *testing.T) {
	convey.Convey("prepare", t, func() {
		s, ok := mirip.Compare(adapter.NewLevenshtein(), testNameSimilar, minimalScore, false, names...)
		convey.So(ok, convey.ShouldBeNil)
		convey.Printf(" %s = %s\n", testNameSimilar, s)
	})
}

func TestFindLevenReverse(t *testing.T) {
	convey.Convey("prepare", t, func() {
		s, ok := mirip.Compare(adapter.NewLevenshtein(), testNameReverse, minimalScore, false, names...)
		convey.So(ok, convey.ShouldBeNil)
		convey.Printf(" %s = %s\n", testNameReverse, s)
	})
}
