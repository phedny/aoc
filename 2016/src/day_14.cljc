(ns day-14 (:require [clojure.string :as str]))

(def salt "ngcjuoqr")

(defn md5
  [s n]
  (if (zero? n)
    s
    (let [algorithm (java.security.MessageDigest/getInstance "MD5")
          raw (.digest algorithm (.getBytes s))]
      (recur (format "%032x" (java.math.BigInteger. 1 raw)) (dec n)))))

(defn groups-of-three
  [s]
  (map second (re-seq #"([0-9a-f])\1\1" s)))

(defn contains-group-of-five?
  [s c]
  (str/index-of s (str c c c c c)))

(defn key-indices
  [n]
  (keep-indexed (fn
                  [index [hash & ts]]
                  (when (not-empty (for [c (take 1 (groups-of-three hash))
                                         t ts
                                         :when (contains-group-of-five? t c)]
                                     true))
                    index))
                (partition 1001 1 (map #(md5 (str salt %) n) (range)))))

(println (nth (key-indices 1) 63))
(println (nth (key-indices 2017) 63))
