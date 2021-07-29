// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"easygoadmin/app/dao/internal"
)

// ConfigDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type ConfigDao struct {
	internal.ConfigDao
}

var (
	// Config is globally public accessible object for table sys_config operations.
	Config = ConfigDao{
		internal.Config,
	}
)

// Fill with you ideas below.