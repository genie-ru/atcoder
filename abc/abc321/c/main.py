import sys
import numpy as np


# コンパイル時
if sys.argv[-1] == 'ONLINE_JUDGE':
    # A問題の答え
    def like321(N: int) -> bool:
        saigo: chr = ':'  # '9' < ':'
        kekka: bool = True
        for suuji in str(N):
            kekka &= suuji < saigo
            saigo = suuji
        return kekka

    # 閉区間で表している、この範囲以外に321-likeな数はない
    hanni: list[tuple[int, int]] =\
     [
        (0, 98_765_432),
        (876_543_210, 987_654_321),
        (9_876_543_210, 9_876_543_210)
     ]
    
    like321_ichiran: np.ndarray = np.empty(0, dtype=np.int64)
    for hajime, owari in hanni:
        like321_konnkai: np.ndarray = \
         np.fromiter(
            filter(
                like321,
                range(hajime, owari + 1)
            ),
            dtype=np.int64
         )
        like321_ichiran = np.append(like321_ichiran, like321_konnkai)

    np.save('./like321_ichiran.npy', like321_ichiran)

# 実行時
else:
    like321_ichiran: np.ndarray = np.load('./like321_ichiran.npy')
    K: int = int(input())
    print(like321_ichiran[K])
