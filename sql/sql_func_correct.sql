IF OBJECT_ID('dbo.GetLastJobId') IS NOT NULL DROP FUNCTION dbo.GetLastJobId;
GO

CREATE FUNCTION dbo.GetLastJobId()
RETURNS TABLE
AS
      RETURN
            SELECT TOP 1
      [JOB_ID]
FROM [PanaCIM].[dbo].[job_history]
where CLOSING_TYPE = '0'
order by END_TIME desc
GO

-- get SETUP_ID
IF OBJECT_ID('dbo.GetLastSetupId') IS NOT NULL DROP FUNCTION dbo.GetLastSetupId;
GO

CREATE FUNCTION dbo.GetLastSetupId()
RETURNS TABLE
    AS
    RETURN
            SELECT [SETUP_ID]
            FROM [PanaCIM].[dbo].[job_products]
--where SETUP_ID = '9536'
            WHERE JOB_ID = (
                SELECT *
                FROM dbo.GetLastJobId())
                --order by [SETUP_ID]
GO

-- get PRODUCT_ID
IF OBJECT_ID('dbo.GetLastProductId') IS NOT NULL DROP FUNCTION dbo.GetLastProductId;
GO

CREATE FUNCTION dbo.GetLastProductId()
    RETURNS TABLE
    AS
    RETURN
SELECT [PRODUCT_ID]
FROM [PanaCIM].[dbo].[product_setup]
WHERE [SETUP_ID] = (
    SELECT *
    FROM dbo.GetLastSetupId()
    )
GO

-- get PATTERN_COMBINATIONS_PER_PANEL
IF OBJECT_ID('dbo.GetQtyPerPanel') IS NOT NULL DROP FUNCTION dbo.GetQtyPerPanel;
GO

CREATE FUNCTION dbo.GetQtyPerPanel()
RETURNS TABLE
    AS
    RETURN
SELECT [PATTERN_COMBINATIONS_PER_PANEL]
FROM [PanaCIM].[dbo].[product_data]
WHERE [PRODUCT_ID] = (
    SELECT *
    FROM dbo.GetLastProductId()
    )
GO
 /*
PRINT dbo.GetLastJobId()*/

/* узнаем кол-во м / з произведенных  */
IF OBJECT_ID('dbo.SUMPattern') IS NOT NULL DROP FUNCTION  dbo.SUMPattern;
GO

CREATE FUNCTION dbo.SUMPattern()
RETURNS TABLE
AS
    RETURN
            SELECT COUNT(DISTINCT PANEL_ID)AS sumPattern
FROM [PanaCIM].[dbo].[panels]
where JOB_ID = (
            SELECT *
FROM dbo.GetLastJobId()) --'5134'
GO
/*
/*  получаем номер NC_VERSION (одной заготовки) - */
/****** Script for SelectTopNRows command from SSMS  ******/
IF OBJECT_ID('dbo.GetNcVersion') IS NOT NULL DROP FUNCTION dbo.GetNcVersion;
GO

CREATE FUNCTION dbo.GetNcVersion()
RETURNS TABLE
      AS
      RETURN
            SELECT top 1
      [NC_VERSION]
FROM [PanaCIM].[dbo].[panels]
where JOB_ID = (SELECT *
FROM dbo.GetLastJobId())
order by NC_VERSION desc
GO
*/
/*
/* получем кол-во плат в заготовке */
IF OBJECT_ID('dbo.CountPCBInPattern') IS NOT NULL DROP FUNCTION dbo.CountPCBInPattern;
GO

CREATE FUNCTION dbo.CountPCBInPattern()
RETURNS TABLE
      AS
      RETURN
            SELECT MAX(PATTERN_NUMBER) AS sumPCBInPattern
FROM [PanaCIM].[dbo].[nc_placement_detail]
where NC_VERSION = (SELECT *
FROM dbo.GetNcVersion()) --'298118'
GO
*/
/* подсчет суммы произведенных плат */
IF OBJECT_ID('dbo.SumProductionPCB') IS NOT NULL DROP FUNCTION dbo.SumProductionPCB;
GO

CREATE FUNCTION dbo.SumProductionPCB()
RETURNS INT
      BEGIN
      RETURN (SELECT *
     -- FROM dbo.SUMPattern()) * (SELECT *
     -- FROM dbo.CountPCBInPattern())
          FROM dbo.SUMPattern()) * (SELECT *
      FROM dbo.GetQtyPerPanel())
END;
GO
/* Создаем представление по результатам спискания компонентов по job_id */
/****** Script for SelectTopNRows command from SSMS  ******/
IF OBJECT_ID('dbo.InfoInstallLastJobId_View', 'V') IS NOT NULL
    DROP VIEW dbo.InfoInstallLastJobId_View
GO

CREATE VIEW dbo.InfoInstallLastJobId_View
AS
      SELECT TOP 10000
            [PanaCIM].[dbo].[Z_CASS_VIEW].[REEL_ID],
            [PanaCIM].[dbo].[reel_data].PART_NO,
            SUM([PanaCIM].[dbo].[Z_CASS_VIEW].PLACE_COUNT) AS PLACE_COUNT,
            SUM([PanaCIM].[dbo].[Z_CASS_VIEW].PICKUP_COUNT) AS PICKUP_COUNT,
            [PanaCIM].[dbo].[REEL_DATA_VIEW].reel_barcode,
            [PanaCIM].[dbo].[reel_data].CURRENT_QUANTITY,
            [PanaCIM].[dbo].[reel_data].QUANTITY AS INITIAL_QUANTITY

      FROM [PanaCIM].[dbo].[Z_CASS_VIEW]

            LEFT JOIN [PanaCIM].[dbo].[REEL_DATA_VIEW]
            ON [PanaCIM].[dbo].[REEL_DATA_VIEW].[reel_id] = [PanaCIM].[dbo].[Z_CASS_VIEW].[REEL_ID]
            LEFT JOIN [PanaCIM].[dbo].[reel_data]
            ON [PanaCIM].[dbo].[REEL_DATA_VIEW].[reel_id] = [PanaCIM].[dbo].[reel_data].REEL_ID
      where [PanaCIM].[dbo].[Z_CASS_VIEW].JOB_ID = (SELECT *
            FROM dbo.GetLastJobId()) AND [PanaCIM].[dbo].[Z_CASS_VIEW].[REEL_ID] IS NOT NULL
      group by [PanaCIM].[dbo].[REEL_DATA_VIEW].reel_barcode, [PanaCIM].[dbo].[Z_CASS_VIEW].REEL_ID, [PanaCIM].[dbo].[reel_data].CURRENT_QUANTITY, [PanaCIM].[dbo].[reel_data].QUANTITY
--order by [PanaCIM].dbo.REEL_DATA_VIEW.reel_barcode desc
GO

/* Функция для расчетов суммарного PLACE_COUNT по reel_id */
IF OBJECT_ID('dbo.SumPLACE_COUNT_ALL_REEL_ID') IS NOT NULL DROP FUNCTION dbo.SumPLACE_COUNT_ALL_REEL_ID;
GO

CREATE FUNCTION dbo.SumPLACE_COUNT_ALL_REEL_ID
(
	@reelid AS INT
--	@lastjobid AS INT
)
RETURNS TABLE
      AS
      RETURN
            SELECT
      --[PanaCIM].[dbo].[Z_CASS_VIEW].[REEL_ID],
      SUM([PanaCIM].[dbo].[Z_CASS_VIEW].PLACE_COUNT) AS PLACE_COUNT
--SUM([PanaCIM].[dbo].[Z_CASS_VIEW].PICKUP_COUNT) AS PICKUP_COUNT
--[PanaCIM].[dbo].[REEL_DATA_VIEW].reel_barcode
FROM [PanaCIM].[dbo].[Z_CASS_VIEW]LEFT JOIN [PanaCIM].[dbo].[REEL_DATA_VIEW]
      ON [PanaCIM].[dbo].[REEL_DATA_VIEW].[reel_id] = [PanaCIM].[dbo].[Z_CASS_VIEW].[REEL_ID]
-- where [PanaCIM].[dbo].[Z_CASS_VIEW].JOB_ID = (SELECT * FROM dbo.GetLastJobId())
Where [PanaCIM].[dbo].[REEL_DATA_VIEW].[reel_id] = @reelid
group by [PanaCIM].[dbo].[REEL_DATA_VIEW].reel_barcode, [PanaCIM].[dbo].[Z_CASS_VIEW].REEL_ID
GO

/* Функция для расчетов суммарного PICKUP_COUNT по reel_id */
CREATE FUNCTION dbo.SumPICKUP_COUNT_ALL_REEL_ID
(
	@reelid AS INT
--	@lastjobid AS INT
)
RETURNS TABLE
      AS
      RETURN
      SELECT
      --[PanaCIM].[dbo].[Z_CASS_VIEW].[REEL_ID],
      --SUM([PanaCIM].[dbo].[Z_CASS_VIEW].PLACE_COUNT) AS PLACE_COUNT
      SUM([PanaCIM].[dbo].[Z_CASS_VIEW].PICKUP_COUNT) AS PICKUP_COUNT_ALL
--[PanaCIM].[dbo].[REEL_DATA_VIEW].reel_barcode
FROM [PanaCIM].[dbo].[Z_CASS_VIEW]LEFT JOIN [PanaCIM].[dbo].[REEL_DATA_VIEW]
      ON [PanaCIM].[dbo].[REEL_DATA_VIEW].[reel_id] = [PanaCIM].[dbo].[Z_CASS_VIEW].[REEL_ID]
-- where [PanaCIM].[dbo].[Z_CASS_VIEW].JOB_ID = (SELECT * FROM dbo.GetLastJobId())
Where [PanaCIM].[dbo].[REEL_DATA_VIEW].[reel_id] = @reelid
group by [PanaCIM].[dbo].[REEL_DATA_VIEW].reel_barcode, [PanaCIM].[dbo].[Z_CASS_VIEW].REEL_ID
GO


SELECT TOP 1000
      [REEL_ID]
        ,[PART_NO]
      , [PLACE_COUNT]
      , [PICKUP_COUNT]
      , [reel_barcode]
      , [CURRENT_QUANTITY]
      , [INITIAL_QUANTITY]
      , (SELECT *
      FROM dbo.SumPLACE_COUNT_ALL_REEL_ID([PanaCIM].[dbo].[InfoInstallLastJobId_View].[REEL_ID])) AS PLACE_COUNT_ALL
      , (SELECT *
      FROM dbo.SumPICKUP_COUNT_ALL_REEL_ID([PanaCIM].[dbo].[InfoInstallLastJobId_View].[REEL_ID])) AS PICKUP_COUNT_ALL
FROM [PanaCIM].[dbo].[InfoInstallLastJobId_View];

SELECT *
FROM dbo.GetLastJobId();
SELECT *
FROM dbo.SUMPattern();
--SELECT *
--FROM dbo.GetNcVersion();
--SELECT *
--FROM dbo.CountPCBInPattern();
PRINT dbo.SumProductionPCB();