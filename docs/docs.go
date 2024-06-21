// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "wanganyang",
            "email": "wayasxxx@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/create": {
            "post": {
                "description": "Create a backup with the given name and collections",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Backup"
                ],
                "summary": "Create backup interface",
                "parameters": [
                    {
                        "type": "string",
                        "description": "request_id",
                        "name": "request_id",
                        "in": "header"
                    },
                    {
                        "description": "CreateBackupRequest JSON",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/backuppb.CreateBackupRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/backuppb.BackupInfoResponse"
                        }
                    }
                }
            }
        },
        "/delete": {
            "delete": {
                "description": "Delete a backup with the given name",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Backup"
                ],
                "summary": "Delete backup interface",
                "parameters": [
                    {
                        "type": "string",
                        "description": "request_id",
                        "name": "request_id",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "backup_name",
                        "name": "backup_name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/backuppb.DeleteBackupResponse"
                        }
                    }
                }
            }
        },
        "/get_backup": {
            "get": {
                "description": "Get the backup with the given name or id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Backup"
                ],
                "summary": "Get backup interface",
                "parameters": [
                    {
                        "type": "string",
                        "description": "request_id",
                        "name": "request_id",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "backup_name",
                        "name": "backup_name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "backup_id",
                        "name": "backup_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/backuppb.BackupInfoResponse"
                        }
                    }
                }
            }
        },
        "/get_restore": {
            "get": {
                "description": "Get restore task state with the given id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Restore"
                ],
                "summary": "Get restore interface",
                "parameters": [
                    {
                        "type": "string",
                        "description": "request_id",
                        "name": "request_id",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/backuppb.RestoreBackupResponse"
                        }
                    }
                }
            }
        },
        "/list": {
            "get": {
                "description": "List all backups in current storage",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Backup"
                ],
                "summary": "List Backups interface",
                "parameters": [
                    {
                        "type": "string",
                        "description": "request_id",
                        "name": "request_id",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "collection_name",
                        "name": "collection_name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/backuppb.ListBackupsResponse"
                        }
                    }
                }
            }
        },
        "/restore": {
            "post": {
                "description": "Submit a request to restore the data from backup",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Restore"
                ],
                "summary": "Restore interface",
                "parameters": [
                    {
                        "type": "string",
                        "description": "request_id",
                        "name": "request_id",
                        "in": "header"
                    },
                    {
                        "description": "RestoreBackupRequest JSON",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/backuppb.RestoreBackupRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/backuppb.RestoreBackupResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "backuppb.BackupInfo": {
            "type": "object",
            "properties": {
                "backup_timestamp": {
                    "description": "backup timestamp",
                    "type": "integer"
                },
                "collection_backups": {
                    "description": "array of collection backup",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/backuppb.CollectionBackupInfo"
                    }
                },
                "end_time": {
                    "type": "integer"
                },
                "errorMessage": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "milvus_version": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "progress": {
                    "type": "integer"
                },
                "size": {
                    "type": "integer"
                },
                "start_time": {
                    "type": "integer"
                },
                "state_code": {
                    "$ref": "#/definitions/backuppb.BackupTaskStateCode"
                }
            }
        },
        "backuppb.BackupInfoResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "response code. 0 means success. others are fail",
                    "allOf": [
                        {
                            "$ref": "#/definitions/backuppb.ResponseCode"
                        }
                    ]
                },
                "data": {
                    "description": "backup info entity",
                    "allOf": [
                        {
                            "$ref": "#/definitions/backuppb.BackupInfo"
                        }
                    ]
                },
                "msg": {
                    "description": "error msg if fail",
                    "type": "string"
                },
                "requestId": {
                    "description": "uuid of the request to response",
                    "type": "string"
                }
            }
        },
        "backuppb.BackupTaskStateCode": {
            "type": "integer",
            "enum": [
                0,
                1,
                2,
                3,
                4
            ],
            "x-enum-varnames": [
                "BackupTaskStateCode_BACKUP_INITIAL",
                "BackupTaskStateCode_BACKUP_EXECUTING",
                "BackupTaskStateCode_BACKUP_SUCCESS",
                "BackupTaskStateCode_BACKUP_FAIL",
                "BackupTaskStateCode_BACKUP_TIMEOUT"
            ]
        },
        "backuppb.Binlog": {
            "type": "object",
            "properties": {
                "entries_num": {
                    "type": "integer"
                },
                "log_path": {
                    "type": "string"
                },
                "log_size": {
                    "type": "integer"
                },
                "timestamp_from": {
                    "type": "integer"
                },
                "timestamp_to": {
                    "type": "integer"
                }
            }
        },
        "backuppb.CollectionBackupInfo": {
            "type": "object",
            "properties": {
                "backup_physical_timestamp": {
                    "description": "physical unix time of backup",
                    "type": "integer"
                },
                "backup_timestamp": {
                    "description": "logical time of backup, used for restore",
                    "type": "integer"
                },
                "channel_checkpoints": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "collection_id": {
                    "type": "integer"
                },
                "collection_name": {
                    "type": "string"
                },
                "consistency_level": {
                    "$ref": "#/definitions/backuppb.ConsistencyLevel"
                },
                "db_name": {
                    "type": "string"
                },
                "end_time": {
                    "type": "integer"
                },
                "errorMessage": {
                    "type": "string"
                },
                "has_index": {
                    "type": "boolean"
                },
                "id": {
                    "type": "string"
                },
                "index_infos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/backuppb.IndexInfo"
                    }
                },
                "l0_segments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/backuppb.SegmentBackupInfo"
                    }
                },
                "load_state": {
                    "type": "string"
                },
                "partition_backups": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/backuppb.PartitionBackupInfo"
                    }
                },
                "progress": {
                    "type": "integer"
                },
                "schema": {
                    "$ref": "#/definitions/backuppb.CollectionSchema"
                },
                "shards_num": {
                    "type": "integer"
                },
                "size": {
                    "type": "integer"
                },
                "start_time": {
                    "type": "integer"
                },
                "state_code": {
                    "$ref": "#/definitions/backuppb.BackupTaskStateCode"
                }
            }
        },
        "backuppb.CollectionSchema": {
            "type": "object",
            "properties": {
                "autoID": {
                    "type": "boolean"
                },
                "description": {
                    "type": "string"
                },
                "enable_dynamic_field": {
                    "type": "boolean"
                },
                "fields": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/backuppb.FieldSchema"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "backuppb.ConsistencyLevel": {
            "type": "integer",
            "enum": [
                0,
                1,
                2,
                3,
                4
            ],
            "x-enum-varnames": [
                "ConsistencyLevel_Strong",
                "ConsistencyLevel_Session",
                "ConsistencyLevel_Bounded",
                "ConsistencyLevel_Eventually",
                "ConsistencyLevel_Customized"
            ]
        },
        "backuppb.CreateBackupRequest": {
            "type": "object",
            "properties": {
                "async": {
                    "description": "async or not",
                    "type": "boolean"
                },
                "backup_name": {
                    "description": "backup name, will generate one if not set",
                    "type": "string"
                },
                "collection_names": {
                    "description": "collection names to backup, empty to backup all",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "db_collections": {
                    "description": "database and collections to backup. A json string. To support database. 2023.7.7",
                    "type": "string"
                },
                "force": {
                    "description": "force backup skip flush, Should make sure data has been stored into disk when using it",
                    "type": "boolean"
                },
                "gc_pause_address": {
                    "description": "gc pause API address",
                    "type": "string"
                },
                "gc_pause_enable": {
                    "description": "if true, stop GC to avoid the data compacted and GCed during backup, use it when the data to backup is very large.",
                    "type": "boolean"
                },
                "gc_pause_seconds": {
                    "description": "gc pause seconds, set it larger than the time cost of backup",
                    "type": "integer"
                },
                "meta_only": {
                    "description": "only backup meta, including collection schema and index info",
                    "type": "boolean"
                },
                "requestId": {
                    "description": "uuid of request, will generate one if not set",
                    "type": "string"
                }
            }
        },
        "backuppb.DataType": {
            "type": "integer",
            "enum": [
                0,
                1,
                2,
                3,
                4,
                5,
                10,
                11,
                20,
                21,
                22,
                23,
                100,
                101,
                102,
                103,
                104
            ],
            "x-enum-varnames": [
                "DataType_None",
                "DataType_Bool",
                "DataType_Int8",
                "DataType_Int16",
                "DataType_Int32",
                "DataType_Int64",
                "DataType_Float",
                "DataType_Double",
                "DataType_String",
                "DataType_VarChar",
                "DataType_Array",
                "DataType_Json",
                "DataType_BinaryVector",
                "DataType_FloatVector",
                "DataType_Float16Vector",
                "DataType_BFloat16Vector",
                "DataType_SparseFloatVector"
            ]
        },
        "backuppb.DeleteBackupResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "response code. 0 means success. others are fail",
                    "allOf": [
                        {
                            "$ref": "#/definitions/backuppb.ResponseCode"
                        }
                    ]
                },
                "msg": {
                    "description": "error msg if fail",
                    "type": "string"
                },
                "requestId": {
                    "description": "uuid of the request to response",
                    "type": "string"
                }
            }
        },
        "backuppb.FieldBinlog": {
            "type": "object",
            "properties": {
                "binlogs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/backuppb.Binlog"
                    }
                },
                "fieldID": {
                    "type": "integer"
                }
            }
        },
        "backuppb.FieldSchema": {
            "type": "object",
            "properties": {
                "autoID": {
                    "type": "boolean"
                },
                "data_type": {
                    "$ref": "#/definitions/backuppb.DataType"
                },
                "default_value": {
                    "$ref": "#/definitions/backuppb.ValueField"
                },
                "description": {
                    "type": "string"
                },
                "element_type": {
                    "$ref": "#/definitions/backuppb.DataType"
                },
                "fieldID": {
                    "type": "integer"
                },
                "index_params": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/backuppb.KeyValuePair"
                    }
                },
                "is_dynamic": {
                    "type": "boolean"
                },
                "is_partition_key": {
                    "type": "boolean"
                },
                "is_primary_key": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "state": {
                    "$ref": "#/definitions/backuppb.FieldState"
                },
                "type_params": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/backuppb.KeyValuePair"
                    }
                }
            }
        },
        "backuppb.FieldState": {
            "type": "integer",
            "enum": [
                0,
                1,
                2,
                3
            ],
            "x-enum-varnames": [
                "FieldState_FieldCreated",
                "FieldState_FieldCreating",
                "FieldState_FieldDropping",
                "FieldState_FieldDropped"
            ]
        },
        "backuppb.IndexInfo": {
            "type": "object",
            "properties": {
                "field_name": {
                    "type": "string"
                },
                "index_name": {
                    "type": "string"
                },
                "index_type": {
                    "type": "string"
                },
                "params": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                }
            }
        },
        "backuppb.KeyValuePair": {
            "type": "object",
            "properties": {
                "key": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "backuppb.ListBackupsResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "response code. 0 means success. others are fail",
                    "allOf": [
                        {
                            "$ref": "#/definitions/backuppb.ResponseCode"
                        }
                    ]
                },
                "data": {
                    "description": "backup info entities",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/backuppb.BackupInfo"
                    }
                },
                "msg": {
                    "description": "error msg if fail",
                    "type": "string"
                },
                "requestId": {
                    "description": "uuid of the request to response",
                    "type": "string"
                }
            }
        },
        "backuppb.PartitionBackupInfo": {
            "type": "object",
            "properties": {
                "collection_id": {
                    "type": "integer"
                },
                "load_state": {
                    "type": "string"
                },
                "partition_id": {
                    "type": "integer"
                },
                "partition_name": {
                    "type": "string"
                },
                "segment_backups": {
                    "description": "array of segment backup",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/backuppb.SegmentBackupInfo"
                    }
                },
                "size": {
                    "type": "integer"
                }
            }
        },
        "backuppb.ResponseCode": {
            "type": "integer",
            "enum": [
                0,
                1,
                2,
                3,
                400,
                404
            ],
            "x-enum-varnames": [
                "ResponseCode_Success",
                "ResponseCode_Not_Support",
                "ResponseCode_No_Permission",
                "ResponseCode_Fail",
                "ResponseCode_Parameter_Error",
                "ResponseCode_Request_Object_Not_Found"
            ]
        },
        "backuppb.RestoreBackupRequest": {
            "type": "object",
            "properties": {
                "async": {
                    "description": "execute asynchronously or not",
                    "type": "boolean"
                },
                "backup_name": {
                    "description": "backup name to restore",
                    "type": "string"
                },
                "bucket_name": {
                    "description": "if bucket_name and path is set. will override bucket/path in config.",
                    "type": "string"
                },
                "collection_names": {
                    "description": "collections to restore",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "collection_renames": {
                    "description": "2, give a map to rename the collections, if not given, use the original name.\ncollection_renames has higher priority than collection_suffix",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "collection_suffix": {
                    "description": "Support two ways to rename the collections while recover\n1, set a suffix",
                    "type": "string"
                },
                "db_collections": {
                    "description": "database and collections to restore. A json string. for example: {\"db1\":[\"collection1\"],\"db2\":[\"collection2\",\"collection3\"]}",
                    "type": "string"
                },
                "dropExistCollection": {
                    "description": "if true, drop existing target collection before create",
                    "type": "boolean"
                },
                "dropExistIndex": {
                    "description": "if true, drop existing index of target collection before create",
                    "type": "boolean"
                },
                "id": {
                    "type": "string"
                },
                "metaOnly": {
                    "description": "if true only restore meta, not restore data",
                    "type": "boolean"
                },
                "path": {
                    "description": "if bucket_name and path is set. will override bucket/path in config.",
                    "type": "string"
                },
                "requestId": {
                    "description": "uuid of request, will generate one if not set",
                    "type": "string"
                },
                "restoreIndex": {
                    "description": "if true restore index info",
                    "type": "boolean"
                },
                "skipCreateCollection": {
                    "description": "if true, will skip collection, use when collection exist, restore index or data",
                    "type": "boolean"
                },
                "useAutoIndex": {
                    "description": "if true use autoindex when restore vector index",
                    "type": "boolean"
                }
            }
        },
        "backuppb.RestoreBackupResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "response code. 0 means success. others are fail",
                    "allOf": [
                        {
                            "$ref": "#/definitions/backuppb.ResponseCode"
                        }
                    ]
                },
                "data": {
                    "description": "restore task info entity",
                    "allOf": [
                        {
                            "$ref": "#/definitions/backuppb.RestoreBackupTask"
                        }
                    ]
                },
                "msg": {
                    "description": "error msg if fail",
                    "type": "string"
                },
                "requestId": {
                    "description": "uuid of the request to response",
                    "type": "string"
                }
            }
        },
        "backuppb.RestoreBackupTask": {
            "type": "object",
            "properties": {
                "collection_restore_tasks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/backuppb.RestoreCollectionTask"
                    }
                },
                "end_time": {
                    "type": "integer"
                },
                "errorMessage": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "progress": {
                    "type": "integer"
                },
                "restored_size": {
                    "type": "integer"
                },
                "start_time": {
                    "type": "integer"
                },
                "state_code": {
                    "$ref": "#/definitions/backuppb.RestoreTaskStateCode"
                },
                "to_restore_size": {
                    "type": "integer"
                }
            }
        },
        "backuppb.RestoreCollectionTask": {
            "type": "object",
            "properties": {
                "coll_backup": {
                    "$ref": "#/definitions/backuppb.CollectionBackupInfo"
                },
                "dropExistCollection": {
                    "description": "if true drop the collections",
                    "type": "boolean"
                },
                "dropExistIndex": {
                    "description": "if true drop index info",
                    "type": "boolean"
                },
                "end_time": {
                    "type": "integer"
                },
                "errorMessage": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "metaOnly": {
                    "description": "if true only restore meta",
                    "type": "boolean"
                },
                "partition_restore_tasks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/backuppb.RestorePartitionTask"
                    }
                },
                "progress": {
                    "type": "integer"
                },
                "restoreIndex": {
                    "description": "if true restore index info",
                    "type": "boolean"
                },
                "restored_size": {
                    "type": "integer"
                },
                "skipCreateCollection": {
                    "description": "if true will skip create collections",
                    "type": "boolean"
                },
                "start_time": {
                    "type": "integer"
                },
                "state_code": {
                    "$ref": "#/definitions/backuppb.RestoreTaskStateCode"
                },
                "target_collection_name": {
                    "type": "string"
                },
                "target_db_name": {
                    "type": "string"
                },
                "to_restore_size": {
                    "type": "integer"
                },
                "useAutoIndex": {
                    "description": "if true use autoindex when restore vector index",
                    "type": "boolean"
                }
            }
        },
        "backuppb.RestorePartitionTask": {
            "type": "object",
            "properties": {
                "end_time": {
                    "type": "integer"
                },
                "errorMessage": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "part_backup": {
                    "$ref": "#/definitions/backuppb.PartitionBackupInfo"
                },
                "progress": {
                    "type": "integer"
                },
                "start_time": {
                    "type": "integer"
                },
                "state_code": {
                    "$ref": "#/definitions/backuppb.RestoreTaskStateCode"
                }
            }
        },
        "backuppb.RestoreTaskStateCode": {
            "type": "integer",
            "enum": [
                0,
                1,
                2,
                3,
                4
            ],
            "x-enum-varnames": [
                "RestoreTaskStateCode_INITIAL",
                "RestoreTaskStateCode_EXECUTING",
                "RestoreTaskStateCode_SUCCESS",
                "RestoreTaskStateCode_FAIL",
                "RestoreTaskStateCode_TIMEOUT"
            ]
        },
        "backuppb.SegmentBackupInfo": {
            "type": "object",
            "properties": {
                "backuped": {
                    "type": "boolean"
                },
                "binlogs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/backuppb.FieldBinlog"
                    }
                },
                "collection_id": {
                    "type": "integer"
                },
                "deltalogs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/backuppb.FieldBinlog"
                    }
                },
                "group_id": {
                    "description": "separate segments into multi groups by size,\nsegments in one group will be copied into one directory during backup\nand will bulkinsert in one call during restore",
                    "type": "integer"
                },
                "is_l0": {
                    "type": "boolean"
                },
                "num_of_rows": {
                    "type": "integer"
                },
                "partition_id": {
                    "type": "integer"
                },
                "segment_id": {
                    "type": "integer"
                },
                "size": {
                    "type": "integer"
                },
                "statslogs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/backuppb.FieldBinlog"
                    }
                }
            }
        },
        "backuppb.ValueField": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "Types that are valid to be assigned to Data:\n\n\t*ValueField_BoolData\n\t*ValueField_IntData\n\t*ValueField_LongData\n\t*ValueField_FloatData\n\t*ValueField_DoubleData\n\t*ValueField_StringData\n\t*ValueField_BytesData"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Milvus Backup Service",
	Description:      "A data backup & restore tool for Milvus",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
