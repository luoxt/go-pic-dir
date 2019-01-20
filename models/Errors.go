package models

import "errors"

var (
    ErrInvalidParameter  = errors.New("方法参数错误")
    
	ErrMemberNoExist             = errors.New("用户不存在")
	ErrMemberExist               = errors.New("用户已存在")
	ErrMemberDisabled            = errors.New("用户被禁用")
	ErrMemberEmailEmpty          = errors.New("用户邮箱不能为空")
	ErrMemberEmailExist          = errors.New("用户邮箱已被使用")
	ErrMemberDescriptionTooLong  = errors.New("用户描述必须小于500字")
	ErrMemberEmailFormatError    = errors.New("邮箱格式不正确")
	ErrMemberPasswordFormatError = errors.New("密码必须在6-50个字符之间")
	ErrMemberAccountFormatError  = errors.New("账号只能由英文字母数字组成，且在3-50个字符")
	ErrMemberRoleError           = errors.New("用户权限不正确")
)
