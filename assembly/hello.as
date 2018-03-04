# データ
      .data
msg:  .asciiz "Hello, World.\n"

# プログラム（テキスト）
    .text
    .globl main
main:                  # 実行を開始する
      li      $v0, 4   # 文字列表示のシステムコール番号（４）
      la      $a0, msg # $a0 レジスタに文字列の番地を入れる
      syscall          # システムコール呼び出し
      jr      $ra      # mainの呼び出し元に戻り、プログラムを終了
# end
