(ns day-9 (:require util))

(def input (first (util/read-input)))

(declare parse-normal)

(defn process-marker-data
  [n cs length count]
  (+ (* length count) (parse-normal n (drop length cs))))

(defn parse-marker-count
  [n [c & cs] length count]
  (if (= c \))
    (process-marker-data n cs length count)
    (recur n cs length (+ (* 10 count) (- (int c) 48)))))

(defn parse-marker-length
  [n [c & cs] length]
  (if (= c \x)
    (parse-marker-count n cs length 0)
    (recur n cs (+ (* 10 length) (- (int c) 48)))))

(defn parse-normal
  [n [c & cs]]
  (cond
    (nil? c) n
    (= c \() (parse-marker-length n cs 0)
    :else (recur (inc n) cs)))

(println (parse-normal 0 input))

(defn process-marker-data
  [n cs length count]
  (+ (* (parse-normal 0 (take length cs)) count) (parse-normal n (drop length cs))))

(println (parse-normal 0 input))
