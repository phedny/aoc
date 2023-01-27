(ns day-25)

(defn triangular
  [n]
  (/ (* n (inc n)) 2))

(defn code-at
  [row column]
  (let [exp (+
              (- (triangular column) (triangular (dec column)))
              (triangular (+ row column -2)))]
    (-> (.modPow (biginteger 252533) (biginteger (dec exp)) (biginteger 33554393))
        (* 20151125)
        (mod 33554393)
        .intValue)))

(println (code-at 2981 3075))
