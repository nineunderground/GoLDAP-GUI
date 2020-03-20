package main

import (
	"os"
)

// ModeNonTLS ..
const ModeNonTLS = "non-TLS"

// ModeTLS ...
const ModeTLS = "TLS"

// ModeSTARTTLS ...
const ModeSTARTTLS = "STARTTLS"

// LdapProtocol ...
const LdapProtocol = "ldap"

// LdapSProtocol ...
const LdapSProtocol = "ldaps"

// LdapPort ...
const LdapPort = "13268" //"389"

// LdapSPort ...
const LdapSPort = "636"

// LdapUser ...
var LdapUser = os.Getenv("LDAP_USERNAME")

// LdapPass ...
var LdapPass = os.Getenv("LDAP_USERPASS")

// LdapBaseDN ...
var LdapBaseDN = os.Getenv("LDAP_BASE_DN")

// Hostname ...
var Hostname = os.Getenv("LDAP_HOSTNAME")

// 389/tcp  open  ldap
// 636/tcp  open  ldapssl

// 3268/tcp open  globalcatLDAP
// 3269/tcp open  globalcatLDAPssl

// AllAttr ... var attributesToReturn = []string{"cn", "name", "mail", "sAMAccountType", "memberOf"}
var AllAttr = []string{"*"}

// MainWindowTitle ...
const MainWindowTitle = "LDAP Browser (ReadOnly) v0.1"

var DefaultFileContent = []string{
	//"# SESSION DETAILS #1",
	"SESSION_1_NAME=Session 1",
	//"# DOMAIN.SERVER",
	"SESSION_1_HOSTNAME=Localhost",
	//"# 389",
	"SESSION_1_PORT=TODO",
	//"# DN=****",
	"SESSION_1_BASE_DN=TODO",
	//"# CN=****",
	"SESSION_1_USER_DN=TODO",
	//"# ****",
	"SESSION_1_PASSWORD=TODO",
	//"# YES|NO",
	"SESSION_1_SSL=NO",
	//"# SESSION DETAILS #2",
	"SESSION_2_NAME=Session 2",
	//"# DOMAIN.SERVER",
	"SESSION_2_HOSTNAME=Localhost",
	//"# 389",
	"SESSION_2_PORT=TODO",
	//"# DN=****",
	"SESSION_2_BASE_DN=TODO",
	//"# CN=****",
	"SESSION_2_USER_DN=TODO",
	//"# ****",
	"SESSION_2_PASSWORD=TODO",
	//"# YES|NO",
	"SESSION_2_SSL=NO",
	//"# SESSION DETAILS #3",
	"SESSION_3_NAME=Session 3",
	//"# DOMAIN.SERVER",
	"SESSION_3_HOSTNAME=Localhost",
	//"# 389",
	"SESSION_3_PORT=TODO",
	//"# DN=****",
	"SESSION_3_BASE_DN=TODO",
	//"# CN=****",
	"SESSION_3_USER_DN=TODO",
	//"# ****",
	"SESSION_3_PASSWORD=TODO",
	//"# YES|NO",
	"SESSION_3_SSL=NO",
	//"# SESSION DETAILS #4",
	"SESSION_4_NAME=Session 4",
	//"# DOMAIN.SERVER",
	"SESSION_4_HOSTNAME=Localhost",
	//"# 389",
	"SESSION_4_PORT=TODO",
	//"# DN=****",
	"SESSION_4_BASE_DN=TODO",
	//"# CN=****",
	"SESSION_4_USER_DN=TODO",
	//"# ****",
	"SESSION_4_PASSWORD=TODO",
	//"# YES|NO",
	"SESSION_4_SSL=NO",
	//"# SESSION DETAILS #5",
	"SESSION_5_NAME=Session 5",
	//"# DOMAIN.SERVER",
	"SESSION_5_HOSTNAME=Localhost",
	//"# 389",
	"SESSION_5_PORT=TODO",
	//"# DN=****",
	"SESSION_5_BASE_DN=TODO",
	//"# CN=****",
	"SESSION_5_USER_DN=TODO",
	//"# ****",
	"SESSION_5_PASSWORD=TODO",
	//"# YES|NO",
	"SESSION_5_SSL=NO"}
