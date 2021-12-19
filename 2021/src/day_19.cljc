(ns day-19 (:require util
                     [clojure.string :refer [split]]
                     [clojure.pprint :refer [pprint]]
                     [clojure.set :refer [intersection]]))

(defn parse-probe-line
  [line]
  (apply vector (map #(Integer/parseInt %) (split line #","))))

(defn parse-scanner-block
  [lines]
  (map parse-probe-line (rest lines)))

(defn parse-input
  [lines]
  (let [block (parse-scanner-block (take-while not-empty lines))
        remaining (drop-while not-empty lines)]
    (if (empty? remaining)
      (list block)
      (cons block (parse-input (rest remaining))))))

(def input
  (parse-input (util/read-input)))

(def rotation-functions
  [(fn [[x y z]] [x y z]) (fn [[x y z]] [x (unchecked-negate y) (unchecked-negate z)])
   (fn [[x y z]] [x z (unchecked-negate y)]) (fn [[x y z]] [x (unchecked-negate z) y])
   (fn [[x y z]] [(unchecked-negate x) y (unchecked-negate z)]) (fn [[x y z]] [(unchecked-negate x) (unchecked-negate y) z])
   (fn [[x y z]] [(unchecked-negate x) z y]) (fn [[x y z]] [(unchecked-negate x) (unchecked-negate z) (unchecked-negate y)])
   (fn [[x y z]] [y x (unchecked-negate z)]) (fn [[x y z]] [y (unchecked-negate x) z])
   (fn [[x y z]] [y z x]) (fn [[x y z]] [y (unchecked-negate z) (unchecked-negate x)])
   (fn [[x y z]] [(unchecked-negate y) x z]) (fn [[x y z]] [(unchecked-negate y) (unchecked-negate x) (unchecked-negate z)])
   (fn [[x y z]] [(unchecked-negate y) z (unchecked-negate x)]) (fn [[x y z]] [(unchecked-negate y) (unchecked-negate z) x])
   (fn [[x y z]] [z x y]) (fn [[x y z]] [z (unchecked-negate x) (unchecked-negate y)])
   (fn [[x y z]] [z y (unchecked-negate x)]) (fn [[x y z]] [z (unchecked-negate y) x])
   (fn [[x y z]] [(unchecked-negate z) x (unchecked-negate y)]) (fn [[x y z]] [(unchecked-negate z) (unchecked-negate x) y])
   (fn [[x y z]] [(unchecked-negate z) y x]) (fn [[x y z]] [(unchecked-negate z) (unchecked-negate y) (unchecked-negate x)])])

(defn create-rotated-scan-results
  [coordinates]
  (reduce #(assoc %1 %2 (map %2 coordinates)) {} rotation-functions))

(defn coordinate-difference
  [c1 c2]
  (apply vector (map - c2 c1)))

(defn coordinate-sum
  [c1 c2]
  (apply vector (map + c1 c2)))

(defn process-scan-results
  [rotation coordinates]
  (->> (for [c1 coordinates
             c2 coordinates
             :when (not= c1 c2)]
         {c1 [{:diff (coordinate-difference c1 c2) :to c2}]})
       (apply merge-with concat)
       (map (fn [[coordinate relatives]] [coordinate (apply vector relatives) rotation]))))

(defn create-transformation-function
  [block1 block2 compose-with]
  (some (fn
           [[coordinate1 relatives1]]
           (let [diffs1 (set (map :diff relatives1))
                 prr-s (mapcat
                         #(apply process-scan-results %)
                         (create-rotated-scan-results block2))]
             (some (fn
                     [[coordinate2 relatives2 rotation]]
                     (let [diffs2 (set (map :diff relatives2))
                           intersecting-diffs (intersection diffs1 diffs2)]
                       (if (> (count intersecting-diffs) 10)
                         (comp compose-with (partial coordinate-sum (coordinate-difference coordinate2 coordinate1)) rotation))))
                   prr-s)))
        (process-scan-results identity block1)))

(defn find-all-transformation-functions
  [transformation-functions]
  (println (count (filter nil? transformation-functions)) "/" (count transformation-functions))
  (if (some nil? transformation-functions)
    (recur (map
             (fn
               [block tf]
               (or
                 tf
                 (->> (map vector input transformation-functions)
                      (filter #(not (nil? (second %))))
                      (some #(create-transformation-function (first %) block (second %))))))
             input transformation-functions))
    transformation-functions))

(def transformation-functions
  (find-all-transformation-functions (cons identity (map (constantly nil) (rest input)))))

(println ">>")

(def all-beacons
  (set (mapcat #(map %2 %1) input transformation-functions)))

(println (count all-beacons))
