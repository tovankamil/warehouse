# jika ada perubahan di sisi branch main dan ingin dipindahkan ke branch lain , lakukan ini

# 1. Simpan semua perubahan Anda sementara (termasuk file baru/untracked)

git stash --include-untracked

# 2. Sekarang Anda bisa berpindah branch dengan aman

git checkout feat/user-role

# 3. Kembalikan perubahan yang tadi disimpan ke dalam branch ini

git stash pop

# 4. Tambahkan ke staging area

git add .

# 5. Lakukan commit

git commit -m "Update user interface dan fungsi repository"

# 6. Push ke GitHub

git push
