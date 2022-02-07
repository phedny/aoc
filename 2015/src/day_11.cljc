(ns day-11)

(defn pass->number
  [pass]
  (reduce #(+ (* 27 %1) (int %2) -96) 0 pass))

(defn number->pass
  ([number]
   (apply str (number->pass number ())))
  ([number suffix]
   (if (zero? number)
     suffix
     (recur (quot number 27) (cons (char (+ (mod number 27) 96)) suffix)))))

(defn contains-no-backtick?
  [pass]
  (not-any? #{\`} pass))

(defn contains-increasing-straight?
  [pass]
  (let [int-pass (map int pass)]
    (some identity (map #(and (= (inc %1) %2) (= (inc %2) %3)) int-pass (next int-pass) (drop 2 int-pass)))))

(defn contains-no-i-o-l?
  [pass]
  (not-any? #{\i \o \l} pass))

(defn contains-two-non-overlapping-pairs?
  [pass]
  (if-let [suffix (drop-while false? (map = pass (next pass)))]
    (and (> (count suffix) 2) (some true? (drop 2 suffix)))))

(def valid-pass?
  (every-pred contains-no-backtick? contains-increasing-straight? contains-no-i-o-l? contains-two-non-overlapping-pairs?))

(def inc-pass (comp number->pass inc pass->number))

(defn inc-valid-pass
  [pass]
  (let [new-pass (inc-pass pass)]
    (if (valid-pass? new-pass)
      new-pass
      (recur new-pass))))

(def pass-0 "cqjxjnds")
(def pass-1 (inc-valid-pass pass-0))
(def pass-2 (inc-valid-pass pass-1))

(println pass-1)
(println pass-2)
