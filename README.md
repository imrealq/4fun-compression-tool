https://codingchallenges.fyi/challenges/challenge-huffman/
https://codingchallenges.substack.com/p/from-the-challenges-huffman

# Các bước thực hiện
1. Read the text and determine the frequency of each character occurring.

+ validate file dựa vào file name
+ đọc file và tính toán số lần xuất hiện

2. Build the binary tree from the frequencies.

tạo binary tree từ kết quả ở bước 1 tính toán số lần xuất hiện
https://opendsa-server.cs.vt.edu/ODSA/Books/CS3/html/Huffman.html

3. Generate the prefix-code table from the tree.

prefix-code table là gì? Prefix-code table là một bảng ánh xạ giữa các ký tự (hoặc ký hiệu) và các chuỗi bit đại diện cho chúng.
https://opendsa-server.cs.vt.edu/ODSA/Books/CS3/html/Huffman.html


4. Encode the text using the code table.

+ hiển thị lựa chọn tên output file
+ viết header vào file output
+ lưu thông tin cần thiết. có 2 cách
    + ghi thông tin tree
    + ghi character frequency
Lưu ý:
    when the header ends and when the compressed data starts.
    header chứa đủ thông tin để tái tạo chính xác cấu trúc dữ liệu cần thiết cho quá trình giải nén

5. Encode the tree - we’ll need to include this in the output file so we can decode it.
+ mã hóa văn bản:
    + Sử dụng prefix-code table đã được tạo ra ở các bước trước.
    + Đọc từng ký tự trong văn bản gốc và thay thế bằng mã tương ứng từ bảng mã.

+ chuyển đổi tiền tố thành chuỗi bit:
    + Mỗi mã trong bảng mã thường là một chuỗi các bit (ví dụ: '0', '10', '110').
    + Kết hợp các chuỗi bit này thành một chuỗi bit dài liên tục.

+ đóng gói bit thành byte:
    + chia chuỗi bit dài thành các nhóm 8 bit (1 byte).
    + chuyển đổi mỗi nhóm 8 bit thành một byte.

+ xử lý bit cuối cùng:
    + mếu số bit cuối cùng không đủ 8 bit, cần có chiến lược để xử lý (ví dụ: thêm các bit 0 vào cuối).
    + ghi thông tin về số bit thêm vào (nếu có) vào phần đầu hoặc cuối file để quá trình giải nén có thể xử lý chính xác.

6. Write the encoded tree and text to an output field
đọc header và encoded lại từ file nén để giải nén

# Lưu ý

1. Nên đọc file từng phần 16, 32 or 64 KB vì file có thể lớn vượt mức RAM
2. Sử dụng bitwise operations.
3. read and write files với binary mode
