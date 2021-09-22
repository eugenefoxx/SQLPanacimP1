
/* поиск последнего закрывшегося job_id */
/****** Script for SelectTopNRows command from SSMS  ******/
SELECT TOP 1 [JOB_ID]
      ,[EQUIPMENT_ID]
      ,[SETUP_ID]
      ,[START_TIME]
      ,[END_TIME]
      ,[CLOSING_TYPE]
      ,[START_OPERATOR_ID]
      ,[END_OPERATOR_ID]
      ,[TFR_REASON]
      ,[LANE_NO]
  FROM [PanaCIM].[dbo].[job_history] where CLOSING_TYPE = '0' order by END_TIME desc 
  
 /* результат - job_id and setup_id */
  
 /* узнаем кол-во м / з произведенных  */
  /****** Script for SelectTopNRows command from SSMS  ******/
SELECT COUNT(DISTINCT PANEL_ID) 
  FROM [PanaCIM].[dbo].[panels] where JOB_ID = '5134'
  
 /* результат1 - 250 */
  
/*  получаем номер NC_VERSION (одной заготовки) - */
  /****** Script for SelectTopNRows command from SSMS  ******/
SELECT top 1 [NC_VERSION]
FROM [PanaCIM].[dbo].[panels] where JOB_ID = '5133' order by NC_VERSION desc

/* результат 298118 */

/* получем кол-во плат в заготовке */
SELECT MAX(PATTERN_NUMBER)
  FROM [PanaCIM].[dbo].[nc_placement_detail] where NC_VERSION = '298118'
/* результат2 - 8 */
*/
2000 плат = результат1 - 250 * результат2 - 8
*/
/* по компонентам -  
узнать уникальные номера панел ид  */
/****** Script for SelectTopNRows command from SSMS  ******/
SELECT DISTINCT PANEL_ID
FROM [PanaCIM].[dbo].[panels] where JOB_ID = '5134' /*order by NC_VERSION desc*/

/* получить список NC_VERSION */
/****** Script for SelectTopNRows command from SSMS  ******/
SELECT DISTINCT NC_VERSION
  FROM [PanaCIM].[dbo].[panels]where JOB_ID = '5134'
  
/*  Получаем список REPORT_ID  */
  /****** Script for SelectTopNRows command from SSMS  ******/
SELECT TOP 1000 [REPORT_ID]
  FROM [PanaCIM].[dbo].[PRODUCTION_REPORTS_NM_NPM_KX_VIEW]where NC_VERSION = '297832' and JOB_ID = '5134' AND PLACE_COUNT != '0'
  
/*  Информация по потреблению компонентов по REPORT_ID  */
  /****** Script for SelectTopNRows command from SSMS  ******/
SELECT [REPORT_ID]
      ,[LOT]
      ,[STAGE]
      ,[TPROD]
      ,[HEAD]
      ,[FADD]
      ,[FSADD]
      ,[FBLKCODE]
      ,[FBLKSERIAL]
      ,[REELID]
      ,[TCNT]
      ,[TMISS]
      ,[RMISS]
      ,[HMISS]
      ,[FMISS]
      ,[MMISS]
      ,[BOARD]
      ,[PARTSNAME]
      ,[PLACE_COUNT]
      ,[TRSMISS]
      ,[TIMESTAMP]
      ,[REEL_ID]
      ,[FEEDER_ID]
      ,[PPIERR_COUNT]
  FROM [PanaCIM].[dbo].[Z_Cass_NPM]where REPORT_ID = '8211776' order by REPORT_ID desc
