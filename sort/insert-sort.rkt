#lang racket
(define (insert-item seq item)
  (cond ((null? seq) (list item))
        ((<= item (car seq)) (cons item seq))
        (else (cons (car seq) 
                    (insert-item (cdr seq) item)))))
       
(define (insert-sort seq)
  (cond ((null? seq) (list))
        (else 
         (insert-item (insert-sort (cdr seq)) (car seq)))))

(insert-sort (list 2 5 3 1 9))
(insert-sort (list 1 1 -1 1))
(insert-sort (list 1))
