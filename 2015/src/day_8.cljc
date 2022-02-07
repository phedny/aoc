(ns day-8 (:require util))

(def input
  (util/read-input))

(defn unquoted-count
  ([s]
   (unquoted-count 0 s))
  ([n s]
   (if (empty? s)
     n
     (case (first s)
       \" (recur n (rest s))
       \\ (case (second s)
            (\" \\) (recur (inc n) (drop 2 s))
            \x (recur (inc n) (drop 4 s)))
       (recur (inc n) (rest s))))))

(defn quoted-count
  ([s]
   (quoted-count 0 s))
  ([n s]
   (if (empty? s)
     (+ 2 n)
     (case (first s)
       (\" \\) (recur (+ 2 n) (rest s))
       (recur (inc n) (rest s))))))

(println (- (apply + (map count input)) (apply + (map unquoted-count input))))
(println (- (apply + (map quoted-count input)) (apply + (map count input))))
