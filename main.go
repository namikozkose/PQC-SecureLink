package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/cloudflare/circl/kem/kyber/kyber768"
)

// =====================================================================
// PROJE: PQC-SecureLink (Post-Quantum Kriptografi Laboratuvarı)
// HEDEF: Kyber-768 (ML-KEM) ile Kuantum Dirençli Anahtar Kapsülleme
// =====================================================================

func main() {
	fmt.Println("[*] Kuantum Sonrası Kriptografi (PQC) Modülü Başlatılıyor...")
	fmt.Println("[*] Algoritma: Kyber-768 (NIST ML-KEM Standardı)\n")

	// Circl kütüphanesinin standart KEM arayüzünü çağırıyoruz
	scheme := kyber768.Scheme()

	// -----------------------------------------------------------------
	// FAZ 1: MERKEZ SOC (ALICE) ANAHTAR ÜRETİMİ
	// -----------------------------------------------------------------
	fmt.Println("--- FAZ 1: MERKEZ SOC (ALICE) ---")
	publicKey, privateKey, err := scheme.GenerateKeyPair()
	if err != nil {
		log.Fatalf("Anahtar üretilemedi: %v", err)
	}

	// Anahtarları byte dizisine (ağda taşınabilir formata) çevirme
	pkBytes, _ := publicKey.MarshalBinary()
	skBytes, _ := privateKey.MarshalBinary()

	fmt.Printf("[+] Alice'in Açık Anahtarı (Public Key) Üretildi. Boyut: %d byte\n", len(pkBytes))
	fmt.Printf("[+] Alice'in Gizli Anahtarı (Private Key) Üretildi. Boyut: %d byte\n", len(skBytes))
	fmt.Println("[>] Açık anahtar, güvensiz internet üzerinden Nükleer Tesise (Bob'a) gönderiliyor...\n")

	// -----------------------------------------------------------------
	// FAZ 2: NÜKLEER TESİS (BOB) ANAHTAR KAPSÜLLEME (ENCAPSULATION)
	// -----------------------------------------------------------------
	fmt.Println("--- FAZ 2: NÜKLEER TESİS (BOB) ---")

	// Bob, Alice'in açık anahtarını kullanarak Kapsülü ve Ortak Sırrı üretir
	ciphertext, bobSharedSecret, err := scheme.Encapsulate(publicKey)
	if err != nil {
		log.Fatalf("Kapsülleme hatası: %v", err)
	}

	fmt.Printf("[+] Bob Ortak Sırrı (Shared Secret) üretti: %x...\n", bobSharedSecret[:8])
	fmt.Printf("[+] Bob Kapsülü (Ciphertext) oluşturdu. Boyut: %d byte\n", len(ciphertext))
	fmt.Println("[>] Kapsül, güvensiz ağ üzerinden Merkez SOC'a (Alice'e) geri gönderiliyor...\n")

	// -----------------------------------------------------------------
	// FAZ 3: MERKEZ SOC (ALICE) KAPSÜL AÇMA (DECAPSULATION)
	// -----------------------------------------------------------------
	fmt.Println("--- FAZ 3: MERKEZ SOC (ALICE) ---")

	// Alice gelen kapsülü kendi gizli anahtarıyla açıp gürültüyü temizler
	aliceSharedSecret, err := scheme.Decapsulate(privateKey, ciphertext)
	if err != nil {
		log.Fatalf("Kapsül açma hatası: %v", err)
	}

	fmt.Printf("[+] Alice Kapsülü çözdü ve Ortak Sırrı elde etti: %x...\n\n", aliceSharedSecret[:8])

	// -----------------------------------------------------------------
	// FAZ 4: GÜVENLİK VE DOĞRULAMA KONTROLÜ
	// -----------------------------------------------------------------
	fmt.Println("--- SONUÇ: DOĞRULAMA ---")
	if bytes.Equal(aliceSharedSecret, bobSharedSecret) {
		fmt.Println("[🛡️] BAŞARILI! Alice ve Bob aynı Ortak Sırra (Shared Secret) ulaştı.")
		fmt.Println("[🛡️] Artık bu 32-byte anahtarı kullanarak AES-256-GCM ile iletişim kurabilirler.")
		fmt.Println("[🛡️] Araya giren bir kuantum bilgisayar o kapsülü (Ciphertext) asla kıramaz!")
	} else {
		fmt.Println("[!] HATA! Anahtarlar eşleşmedi, araya müdahale edildi veya gürültü çözülemedi.")
	}
}
