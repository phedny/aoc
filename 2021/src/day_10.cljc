(ns day-10 (:require util))

(def input
  (util/read-input seq))

(def char-pairs
  {\( \) \[ \] \{ \} \< \>})

(defn traverse
  [input stack]
  (cond
    (empty? input) nil
    (contains? char-pairs (first input)) (recur (rest input) (cons (first input) stack))
    (= (first input) (char-pairs (first stack))) (recur (rest input) (rest stack))
    :else (first input)))

(def wrong-char-score
  {\) 3 \] 57 \} 1197 \> 25137 nil 0})

(def result-part-a
  (apply + (map #(wrong-char-score (traverse % '())) input)))

(print result-part-a)
