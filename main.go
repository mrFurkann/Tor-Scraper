package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/net/proxy"
)

func main() {

	logDosyasi, err := os.OpenFile("scan_report.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Log dosyası oluşturulamadı:", err)
		return
	}
	defer logDosyasi.Close()

	if _, err := os.Stat("output_data"); os.IsNotExist(err) {
		os.Mkdir("output_data", 0755)
	}

	fmt.Println("--- TOR TARAYICI BAŞLIYOR ---")

	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:9150", nil, proxy.Direct)
	if err != nil {
		fmt.Println("HATA: Tor proxy'sine bağlanılamadı. Tor açık mı?")
		return
	}

	transport := &http.Transport{
		Dial: dialer.Dial,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second,
	}

	fmt.Println("IP kontrolü yapılıyor...")
	resp, err := client.Get("http://check.torproject.org/api/ip")
	if err == nil {

		ipBody, _ := io.ReadAll(resp.Body)
		fmt.Println("Gizli IP Adresimiz:", strings.TrimSpace(string(ipBody)))
		resp.Body.Close()
	} else {
		fmt.Println("IP kontrolü başarısız (ama devam ediyoruz):", err)
	}

	dosya, err := os.Open("targets.yaml")
	if err != nil {
		fmt.Println("targets.yaml dosyası bulunamadı! Oluşturmayı unutma.")
		return
	}
	defer dosya.Close()

	scanner := bufio.NewScanner(dosya)

	for scanner.Scan() {
		url := scanner.Text()
		url = strings.TrimSpace(url)

		if url == "" {
			continue
		}

		if !strings.HasPrefix(url, "http") {
			url = "http://" + url
		}

		mesaj := fmt.Sprintf("[TARANIYOR] %s ...\n", url)
		fmt.Print(mesaj)
		logDosyasi.WriteString(mesaj)

		cevap, err := client.Get(url)
		if err != nil {
			hataMesaji := fmt.Sprintf("[HATA] Erişim yok: %s\n", url)
			fmt.Print(hataMesaji)
			logDosyasi.WriteString(hataMesaji)
			continue
		}

		htmlVerisi, err := io.ReadAll(cevap.Body)
		cevap.Body.Close()

		if err != nil {
			fmt.Println("Veri okunurken hata oluştu.")
			continue
		}

		dosyaAdi := strings.ReplaceAll(url, "http://", "")
		dosyaAdi = strings.ReplaceAll(dosyaAdi, ".onion", "")
		dosyaAdi = strings.ReplaceAll(dosyaAdi, "/", "_")
		kayitYolu := fmt.Sprintf("output_data/%s_%d.html", dosyaAdi, time.Now().Unix())

		err = os.WriteFile(kayitYolu, htmlVerisi, 0644)

		if err != nil {
			fmt.Println("Dosya yazılamadı:", err)
		} else {
			basariMesaji := fmt.Sprintf("[BAŞARILI] Kaydedildi: %s\n", kayitYolu)
			fmt.Print(basariMesaji)
			logDosyasi.WriteString(basariMesaji)
		}
	}

	fmt.Println("--- TARAMA BİTTİ ---")
}
