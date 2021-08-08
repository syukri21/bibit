SELECT
    "USER"."ID",
    "USER"."UserName",
    ( CASE "USER"."Parent" WHEN 0 THEN NULL ELSE ( SELECT "U"."UserName" FROM "USER" "U" WHERE "ID" = "USER"."ID" ) END ) AS "ParentUserName"
FROM
    "USER"

-- OR

SELECT "c"."ID",
       "c"."UserName",
       "u"."UserName" as "ParentName"
FROM "USER" "c"
LEFT JOIN "USER" "u" ON "u"."Parent" = "c"."ID"