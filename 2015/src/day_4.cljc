(ns day-4)

(defn md5
  [s]
  (let [algorithm (java.security.MessageDigest/getInstance "MD5")
        raw (.digest algorithm (.getBytes s))]
    (format "%032x" (java.math.BigInteger. 1 raw))))

(defn find-coin
  [prefix secret n]
  (let [coin (md5 (str secret n))]
    (if (= (repeat prefix \0) (take prefix coin))
      n
      (recur prefix secret (inc n)))))

(println (find-coin 5 "iwrupvqb" 1))
(println (find-coin 6 "iwrupvqb" 1))
