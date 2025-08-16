# Lastsafe-PQC

Lastsafe-PQC เป็นเครื่องมือสำรองข้อมูลที่ต่อยอดจาก Restic และ Rclone โดยเพิ่มชั้นการเข้ารหัส Post-Quantum Cryptography (PQC) เพื่อให้มั่นใจว่าข้อมูลของคุณจะปลอดภัยแม้ในยุคคอมพิวเตอร์ควันตัม และมาพร้อมกับ GUI ที่ใช้งานง่าย

## คุณสมบัติ
- **สำรองข้อมูลด้วย Restic**: สร้างและจัดการ snapshot ของไฟล์ในระบบของคุณอย่างมีประสิธิภาพ
- **ซิงค์กับ Rclone**: ส่งข้อมูลที่สำรองไปยังผู้ให้บริการ cloud storage ต่าง ๆ (เช่น S3, Google Drive) ด้วย Rclone
- **เข้ารหัส Post-Quantum**: ใช้อัลกอริทึม Kyber KEM + AES-GCM แบบไฮบริด เพื่อป้องกันการโจมตีของคอมพิวเตอร์ควันตัม
- **GUI ที่เรียบง่าย**: อินเทอร์เฟสที่พัฒนาด้วย Fyne ช่วยให้สำรอง/กู้คืน/ซิงค์ข้อมูลได้สะดวก
- **รองรับหลายระบบปฏิบัติการ**: Windows, Linux และ macOS

## การติดตั้ง

1. ติดตั้ง [Go](https://golang.org/) รุ่น 1.21 หรือใหม่กว่า และติดตั้ง Restic และ Rclone ตามเว็บไซต์ไซต์ทางการ
2. คลอนหรือดาวน์โค้ดจาก GitHub:
   ```bash
   git clone https://github.com/nontawattalk/lastsafe-pqc.git
   cd lastsafe-pqc
   go mod download
   ```
3. สร้างไบนารี หรือรันโปรแกรม:
   ```bash
   go build ./cmd/lastsafe
   ./lastsafe
   ```
   หรือใช้ `go run` เพื่อทดสอบทันที่:
   ```bash
   go run ./cmd/lastsafe
   ```

## วิธีใช้งาน
- **ใช้งานผ่าน GUI**: เปิดโปรแกรม Lastsafe แล้วคลิกปุ่ม `Backup Now` เพื่อสำรองข้อมูล, `Restore` เพื่อกู้คืน และ `Sync to Cloud` เพื่อซิงค์ข้อมูลไปยังคลาวด
- **กำหนดค่า**: ในหน้าตั้งค่า สามารถ เลือกโฟลเดอร์ต้นทาง, ทีเก็บ repository, บัญชี rclone remote และเปิดการเข้ารหัส PQC ได้
- **สำรองข้อมูลผ่าน command line**: หากต้องการใช้ Restic หรือ Rclone โดยตรง ตัวโปรแกรมของเราจะเรียก restic เช่น:
   ```bash
   restic -r /path/to/repo backup /data
   ```
- **ซิงค์ด้วย Rclone**: ตัวอย่างการซิงค์โฟลเดอรสำรองไปยัง S3:
   ```bash
   rclone sync /path/to/repo remote:s3bucket
   ```

## License
This project is licensed under the GPL-3.0 License. See the [LICENSE](LICENSE) file for details.
