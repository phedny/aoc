(ns day-5 (:require util))

(defn md5
  [s]
  (let [algorithm (java.security.MessageDigest/getInstance "MD5")
        raw (.digest algorithm (.getBytes s))]
    (format "%032x" (java.math.BigInteger. 1 raw))))

(defn find-good-hash
  [prefix secret n]
  (let [coin (md5 (str secret n))]
    (if (= (repeat prefix \0) (take prefix coin))
      n
      (recur prefix secret (inc n)))))

(defn find-password-a
  [prefix secret length]
  (->> (iterate #(find-good-hash prefix secret (inc %)) 0)
       next
       (take length)
       (map #(get (md5 (str secret %)) prefix))
       (apply str)))

(println (find-password-a 5 "cxdnnyjw" 8))

(defn find-password-b
  [prefix secret length]
  (loop [n 0
         password {}]
    (if (= (count password) length)
      (apply str (map password (range length)))
      (let [n (find-good-hash prefix secret (inc n))
            hash (md5 (str secret n))
            pos (- (int (get hash prefix)) 48)
            chr (get hash (inc prefix))]
        (if (and (< pos length) (not (contains? password pos)))
          (recur n (assoc password pos chr))
          (recur n password))))))

(println (find-password-b 5 "cxdnnyjw" 8))
