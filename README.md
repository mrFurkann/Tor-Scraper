# Go Onion Scraper (Tor Ağı Veri Toplama Aracı)

Bu proje, **Go (Golang)** programlama dili kullanılarak geliştirilmiş, **Tor ağı** üzerindeki `.onion` uzantılı web sitelerini tarayan ve veri toplayan bir araçtır.

Uygulama, yerel bir **SOCKS5 Proxy** üzerinden trafiği yönlendirerek, IP adresi gizliliğini korur ve hedef sitelerden HTML verilerini çeker. Ağ güvenliği, web scraping ve proxy yönetimi konularında pratik amaçlı geliştirilmiştir.

## Özellikler

* **Tor Entegrasyonu:** Tüm HTTP trafiğini yerel Tor SOCKS5 proxy (127.0.0.1:9050/9150) üzerinden geçirir.
* **IP Gizliliği Kontrolü:** Tarama başlamadan önce `check.torproject.org` üzerinden IP kontrolü yapar.
* **Hata Yönetimi:** Erişilemeyen veya zaman aşımına (timeout) uğrayan siteler programı durdurmaz; hata loglanır ve tarama devam eder.
* **Dosya Kaydı:** Başarılı isteklerden dönen HTML verileri, zaman damgasıyla birlikte `output_data` klasörüne kaydedilir.
* **Loglama:** Tüm işlemler ve hatalar `scan_report.log` dosyasına detaylıca yazılır.

## Gereksinimler

* **Go:** 1.18 veya üzeri sürüm.
* **Tor Browser** veya **Tor Servisi:** Arka planda çalışıyor olmalıdır.
* **Kütüphane:** `golang.org/x/net/proxy`

## Kurulum

1.  Projeyi klonlayın:
    ```bash
    git clone [https://github.com/KULLANICIADIN/PROJE-ADIN.git](https://github.com/KULLANICIADIN/PROJE-ADIN.git)
    cd PROJE-ADIN
    ```

2.  Gerekli bağımlılığı indirin:
    ```bash
    go get golang.org/x/net/proxy
    ```

3.  Hedef listesini oluşturun:
    Proje dizininde `targets.yaml` adında bir dosya oluşturun ve taramak istediğiniz onion adreslerini alt alta ekleyin.

## Kullanım

1.  Tor servisini veya Tor Browser'ı başlatın.
2.  Uygulamayı çalıştırın:
    ```bash
    go run main.go
    ```
3.  Sonuçlar `output_data/` klasörüne, loglar ise `scan_report.log` dosyasına kaydedilecektir.

## Yasal Uyarı

Bu yazılım tamamen **eğitim ve öğrenim amaçlı** geliştirilmiştir. Kötü niyetli kullanımlardan veya yasa dışı sitelere erişimden doğabilecek sorumluluk kullanıcıya aittir.

---
