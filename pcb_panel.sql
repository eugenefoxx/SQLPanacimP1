/****** Script for SelectTopNRows command from SSMS  ******/
SELECT TOP 1000 [PANEL_PATTERN_ID]
      ,[BARCODE]
      ,[MODEL_NO]
      ,[SERIAL_NO]
      ,[LAST_MODIFIED_TIME]
      ,[OPERATOR_ID]
FROM [PanaCIM].[dbo].[panel_details]
WHERE BARCODE = '100473A061822500'

/****** Script for SelectTopNRows command from SSMS  ******/
SELECT TOP 1000 [PANEL_PLACEMENT_ID]
      ,[REEL_ID]
      ,[NC_PLACEMENT_ID]
      ,[PATTERN_NO]
      ,[Z_NUM]
      ,[PU_NUM]
      ,[PART_NO]
      ,[CUSTOM_AREA1]
      ,[CUSTOM_AREA2]
      ,[CUSTOM_AREA3]
      ,[CUSTOM_AREA4]
      ,[REF_DESIGNATOR]
      ,[PATTERN_IDNUM]
      ,[PATTERN_BARCODE]
      ,[PATTERN_DESIGNATOR]
      ,[PREPICKUP_LOT]
      ,[PREPICKUP_STS]
      ,[NCADD]
      ,[NHADD]
      ,[NBLKSERIAL]
      ,[PLACEMENT_ORDER]
      ,[FEEDER_BC]
FROM [PanaCIM].[dbo].[panel_placement_details]

/****** Script for SelectTopNRows command from SSMS  ******/
SELECT TOP 1000 [SERIAL_NO]
      ,[PROD_MODEL_NO]
      ,[PANEL_ID]
      ,[PATTERN_ID]
      ,[BARCODE]
      ,[SETUP_ID]
      ,[TOP_BOTTOM]
      ,[TIMESTAMP]
      ,[IMPORT_FLAG]
FROM [PanaCIM].[dbo].[tracking_data]

-- метод узнать кол-во типов м/з в продукте в колонке PATTERN_TYPES_PER_PANEL
SELECT TOP 1000 [PRODUCT_ID]
      ,[PRODUCT_NAME]
      ,[DOS_PRODUCT_NAME]
      ,[PATTERNS_PER_PANEL]
      ,[PANEL_WIDTH]
      ,[PANEL_LENGTH]
      ,[PANEL_THICKNESS]
      ,[CAMERA_XAXIS_TOP]
      ,[CAMERA_YAXIS_TOP]
      ,[CAMERA_XAXIS_BOTTOM]
      ,[CAMERA_YAXIS_BOTTOM]
      ,[TOOLING_PIN_DISTANCE]
      ,[BARCODES_PER_PANEL]
      ,[PRODUCT_VALID_FLAG]
      ,[TOOLING_PIN]
      ,[CONVEYOR_SPEED]
      ,[USE_BRD_FILE]
      ,[BASE_PRODUCT_ID]
      ,[PATTERN_COMBINATIONS_PER_PANEL]
      ,[PATTERN_TYPES_PER_PANEL]
FROM [PanaCIM].[dbo].[product_data]
    -- where PRODUCT_ID = '3227'
where PRODUCT_ID = '2868'

-- узнать PRODUCT_ID через SETUP_ID
/****** Script for SelectTopNRows command from SSMS  ******/
SELECT TOP 1000 [PRODUCT_ID]
      ,[ROUTE_ID]
      ,[MIX_NAME]
      ,[SETUP_ID]
      ,[LDF_FILE_NAME]
      ,[MACHINE_FILE_NAME]
      ,[SETUP_VALID_FLAG]
      ,[LAST_MODIFIED_TIME]
      ,[DOS_FILE_NAME]
      ,[MODEL_STRING]
      ,[TOP_BOTTOM]
      ,[PT_GROUP_NAME]
      ,[PT_LOT_NAME]
      ,[PT_MC_FILE_NAME]
      ,[PT_DOWNLOADED_FLAG]
      ,[PT_NEEDS_DOWNLOAD]
      ,[SUB_PARTS_FLAG]
      ,[BARCODE_SIDE]
      ,[CYCLE_TIME]
      ,[IMPORT_SOURCE]
      ,[MODIFIED_IMPORT_SOURCE]
      ,[THEORETICAL_XOVER_TIME]
      ,[PUBLISH_MODE]
      ,[PCB_NAME]
      ,[MASTER_MJS_ID]
      ,[LED_VALID_FLAG]
      ,[DGS_PPD_VALID_FLAG]
      ,[ACTIVE_MODEL_STRING]
      ,[REGISTERED_PCB_NAME]
FROM [PanaCIM].[dbo].[product_setup]
    --where MACHINE_FILE_NAME = 'ID9536'
order by LAST_MODIFIED_TIME desc