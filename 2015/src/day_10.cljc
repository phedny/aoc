(ns day-10)

(defn look-and-say
  ([input]
   (apply str (look-and-say input nil 0 [])))
  ([input last-digit last-digit-count prefix]
   (cond
     (empty? input) (if (zero? last-digit-count) prefix (conj prefix last-digit-count last-digit))
     (= (first input) last-digit) (recur (rest input) last-digit (inc last-digit-count) prefix)
     (zero? last-digit-count) (recur (rest input) (first input) 1 prefix)
     :else (recur (rest input) (first input) 1 (conj prefix last-digit-count last-digit)))))

(defn n-times-look-and-say
  [n init]
  (reduce (fn [input _] (look-and-say input)) init (range n)))

(println (count (n-times-look-and-say 40 "1113222113")))
(println (count (n-times-look-and-say 50 "1113222113")))
