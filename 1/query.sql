SELECT
    "USER"."ID",
    "USER"."UserName",
    ( CASE "USER"."Parent" WHEN 0 THEN NULL ELSE ( SELECT "U"."UserName" FROM "USER" "U" WHERE "ID" = "USER"."ID" ) END ) AS "ParentUserName"
FROM
    "USER"