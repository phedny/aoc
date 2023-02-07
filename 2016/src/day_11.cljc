(ns day-11 (:require util))

(def input (map set (util/read-input #(map (fn [[_ element generator?]] [(keyword element) (if generator? :generator :microchip)])
                                           (re-seq #"a (\w+)( generator)?" %)))))

(defn indices-of
  [pred coll]
  (keep-indexed #(when (pred %2) %1) coll))

(def transformed-input
  (let [items (mapv (fn [item]
                      {:generator (first (indices-of #(contains? % [item :generator]) input))
                       :microchip (first (indices-of #(contains? % [item :microchip]) input))})
                   (distinct (mapcat #(map first %) input)))]
    {:items items :elevator-at 0}))

(defn indices
  [pred coll]
  (->> coll
       (keep-indexed #(when (pred %2) [%1 %2]))
       (group-by second)
       (map #(map first (second %)))))

(defn move-one-item
  [{:keys [elevator-at items]} item-type]
  (let [movable-items (map first (indices #(= (item-type %) elevator-at) items))]
    (mapcat (fn [item]
              [{:elevator-at (inc elevator-at) :items (update-in items [item item-type] inc)}
               {:elevator-at (dec elevator-at) :items (update-in items [item item-type] dec)}])
            movable-items)))

(defn move-two-items
  [{:keys [elevator-at items]} item-type-1 item-type-2]
  (let [movable-items-1 (indices #(= (item-type-1 %) elevator-at) items)
        movable-items-2 (indices #(= (item-type-2 %) elevator-at) items)]
    (->> (for [item-1 movable-items-1
               item-2 movable-items-2
               :let [item-1 (first item-1)
                     item-2 (if (and (= item-type-1 item-type-2) (= item-1 (first item-2)))
                              (second item-2)
                              (first item-2))]
               :when item-2]
           [{:elevator-at (inc elevator-at) :items (-> items (update-in [item-1 item-type-1] inc) (update-in [item-2 item-type-2] inc))}
            {:elevator-at (dec elevator-at) :items (-> items (update-in [item-1 item-type-1] dec) (update-in [item-2 item-type-2] dec))}])
        (mapcat identity)
        distinct)))

(defn valid-layout?
  [{:keys [elevator-at items]}]
  (and (<= 0 elevator-at 3)
       (empty? (for [item-1 items
                     item-2 items
                     :when (and (= (:microchip item-1) (:generator item-2)) (not= (:microchip item-1) (:generator item-1)))]
                 [item-1 item-2]))))

(defn can-assemble?
  [{:keys [items]}]
  (every? #(= 3 (:microchip %) (:generator %)) items))

(defn valid-moves
  [layout]
  (->> (concat (move-two-items layout :microchip :generator)
               (move-two-items layout :microchip :microchip)
               (move-two-items layout :generator :generator)
               (move-one-item layout :microchip)
               (move-one-item layout :generator))
       (filter valid-layout?)))

(defn calc-estimate
  [{:keys [elevator-at items]}]
  (let [items-at (mapcat vals items)]
    (- (reduce + (map #(* 2 (- 3 %)) items-at))
       (* (if (= (count (filter #(= elevator-at %) items-at)) 1) 1 3) (- 3 elevator-at)))))

(defn compare-estimate-cost-layout
  [[estimate-a cost-a layout-a] [estimate-b cost-b layout-b]]
  (compare [estimate-a cost-a (.hashCode layout-a)] [estimate-b cost-b (.hashCode layout-b)]))

(defn find-fewest-number-of-steps
  [layout]
  (loop [todo (sorted-set-by compare-estimate-cost-layout [(calc-estimate layout) 0 layout])
         seen {}
         best-cost Integer/MAX_VALUE]
    (let [[estimate cost layout :as item] (first todo)
          todo (disj todo item)]
      (cond
        (<= best-cost estimate) best-cost
        (can-assemble? layout) (recur todo seen cost)
        :else (recur (->> layout
                          valid-moves
                          (remove #(<= (get seen % Integer/MAX_VALUE) cost))
                          (map #(vector (+ 1 cost (calc-estimate %)) (inc cost) %))
                          (reduce conj todo))
                     (assoc seen layout cost) best-cost)))))

(println (find-fewest-number-of-steps transformed-input))
(println (find-fewest-number-of-steps (update transformed-input :items conj {:microchip 0 :generator 0} {:microchip 0 :generator 0})))
