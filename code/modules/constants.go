package api

import "os"

const ModeNonTLS = "non-TLS"
const ModeTLS = "TLS"
const ModeSTARTTLS = "STARTTLS"
const LdapProtocol = "ldap"
const LdapSProtocol = "ldaps"
const LdapPort = "13268" //"389"
const LdapSPort = "636"

var LdapUser = os.Getenv("LDAP_USERNAME")
var LdapPass = os.Getenv("LDAP_USERPASS")
var LdapBaseDN = os.Getenv("LDAP_BASE_DN")
var Hostname = os.Getenv("LDAP_HOSTNAME")

// 389/tcp  open  ldap
// 636/tcp  open  ldapssl

// 3268/tcp open  globalcatLDAP
// 3269/tcp open  globalcatLDAPssl

// var attributesToReturn = []string{"cn", "name", "mail", "sAMAccountType", "memberOf"}
var AllAttr = []string{"*"}

// GUI VALUES

const MainWindowTitle = "LDAP Browser Editor v0.1"
