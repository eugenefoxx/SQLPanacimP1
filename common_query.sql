/****** Script for SelectTopNRows command from SSMS  ******/
-- посмотреть пару job_id и setup_id
SELECT TOP 5000 [JOB_ID]
      ,[SETUP_ID]
FROM [PanaCIM].[dbo].[job_products]
order by [SETUP_ID]

-- посмотреть список ордеров по изделиям
SELECT TOP 100000 [WORK_ORDER_ID]
      ,[WORK_ORDER_NAME]
      ,[LOT_SIZE]
      ,[JOB_ID]
      ,[MASTER_WORK_ORDER_ID]
      ,[COMMENTS]
FROM [PanaCIM].[dbo].[work_orders]
order by [JOB_ID]

SELECT [PCB_NAME]
FROM [PanaCIM].[dbo].[product_setup]
WHERE PRODUCT_ID = '3410'
group by [PCB_NAME]
