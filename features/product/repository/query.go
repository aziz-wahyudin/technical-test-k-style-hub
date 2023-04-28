package repository

/*
sql for handling this task:
SELECT m.username, m.gender, m.skintype, m.skincolor, r.desc_review, SUM(l.likes) AS likes, SUM(l.dislikes) AS dislikes
FROM PRODUCT p
INNER JOIN REVIEW_PRODUCT r ON r.id_product = p.id_product
INNER JOIN MEMBER m ON m.id_member = r.id_member
LEFT JOIN LIKE_REVIEW l ON l.id_review = r.id_review
WHERE p.id_product = <your_product_id_here>
GROUP BY m.username, m.gender, m.skintype, m.skincolor, r.desc_review;

*/
