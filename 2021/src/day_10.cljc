(ns day-10 (:require util))

(def input
  (util/read-input seq))

(def char-pairs
  {\( \) \[ \] \{ \} \< \>})

(defn traverse
  [input stack]
  (cond
    (empty? input) stack
    (contains? char-pairs (first input)) (recur (rest input) (cons (first input) stack))
    (= (first input) (char-pairs (first stack))) (recur (rest input) (rest stack))
    :else (first input)))

(def wrong-char-score
  {\) 3 \] 57 \} 1197 \> 25137})

(def result-part-a
  (apply + (filter (complement nil?) (map #(wrong-char-score (traverse % '())) input))))

(println result-part-a)

(defn points-to-close
  [xs]
  (reduce #(+ (* 5 %1) ({\( 1 \[ 2 \{ 3 \< 4} %2)) 0 xs))

(defn median
  [xs]
  (first (drop (quot (count xs) 2) xs)))

(def result-part-b
  (->> input (map #(traverse % '())) (filter seq?) (map points-to-close) sort median))

(println result-part-b)
