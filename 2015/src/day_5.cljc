(ns day-5 (:require util))

(def input
  (util/read-input))

(defn at-least-three-vowels
  [input]
  (<= 3 (count (filter #(contains? (set "aeiou") %) input))))

(defn at-least-once-two-identical-letters
  [input]
  (< 0 (count (filter #(= (count (set %)) 1) (partition 2 1 input)))))

(defn no-banned-strings
  [input]
  (not (some #(not (neg? (.indexOf input %))) ["ab" "cd" "pq" "xy"])))

(defn contains-pair-of-two-letters
  [input]
  (cond
    (< (count input) 4) false
    (neg? (.indexOf (subs input 2) (subs input 0 2))) (recur (subs input 1))
    :else true))

(defn contains-repeating-letter-with-letter-between
  [input]
  (some (fn [[a _ b]] (= a b)) (partition 3 1 input)))

(defn compute-result
  [& predicates]
  (count (filter (apply every-pred predicates) input)))

(println (compute-result at-least-three-vowels at-least-once-two-identical-letters no-banned-strings))
(println (compute-result contains-pair-of-two-letters contains-repeating-letter-with-letter-between))
