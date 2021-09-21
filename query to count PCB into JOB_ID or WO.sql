/* По ворк-ордеру определяем job_id */
/****** Script for SelectTopNRows command from SSMS  ******/
SELECT TOP 1000 [WORK_ORDER_ID]
      ,[WORK_ORDER_NAME]
      ,[LOT_SIZE]
      ,[JOB_ID]
      ,[MASTER_WORK_ORDER_ID]
      ,[COMMENTS]
  FROM [PanaCIM].[dbo].[work_orders]where WORK_ORDER_NAME = 'No5897_TM11_pri'
  /*Результат*.
  WORK_ORDER_ID,WORK_ORDER_NAME,LOT_SIZE,JOB_ID,MASTER_WORK_ORDER_ID,COMMENTS
6621,No5897_TM11_pri,3000,5131,-1,NULL

/* Узнаем job_id и по нему проводим расчет произведенных м/з*/
/****** Script for SelectTopNRows command from SSMS  ******/
SELECT COUNT(DISTINCT PANEL_ID) /*TOP 1000 [PANEL_ID]
      ,[EQUIPMENT_ID]
      ,[NC_VERSION]
      ,[START_TIME]
      ,[END_TIME]
      ,[PANEL_EQUIPMENT_ID]
      ,[PANEL_SOURCE]
      ,[PANEL_TRACE]
      ,[STAGE_NO]
      ,[LANE_NO]
      ,[JOB_ID]
      ,[SETUP_ID]
      ,[TRX_PRODUCT_ID]
      ,[HUMIDITY]
      ,[TEMPERATURE]
      ,[SETUP_AUDIT_HISTORY_ID]
      ,[AUDIT_STATUS]*/
  FROM [PanaCIM].[dbo].[panels] where JOB_ID = '5131'
  Результат - 750 
  
  // Узнать кол-во плат в заготовке 
  /****** Script for SelectTopNRows command from SSMS  ******/
SELECT MAX(PATTERN_NUMBER)/*COUNT(PATTERN_NUMBER)*/ /*TOP 1000 [NC_PLACEMENT_ID]
      ,[IDNUM]
      ,[REF_DESIGNATOR]
      ,[PART_NAME]
      ,[PATTERN_NUMBER]
      ,[NC_VERSION]
      ,[PATTERN_IDNUM]
      ,[PU_DISPLAY]*/
  FROM [PanaCIM].[dbo].[nc_placement_detail]where NC_VERSION = '298113' /*AND PATTERN_NUMBER = '1'*/
