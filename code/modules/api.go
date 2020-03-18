package api

import (
	"crypto/tls"
	"fmt"

	ldap "github.com/go-ldap/ldap/v3"
)

func Connect(mode string) *ldap.Conn {
	if mode == ModeNonTLS {
		ldapServer := LdapProtocol + "://" + Hostname + ":" + LdapPort
		return ConnectUnsecureDialURL(ldapServer)
		fmt.Println("Connect ModeNonTLS...") // TESTED OK
	} else if mode == ModeTLS {
		ldapServer := LdapSProtocol + "://" + Hostname + ":" + LdapSPort
		return ConnectSecureDialURL(ldapServer)
		fmt.Println("Connect ModeTLS...") // NOT TESTED
	} else if mode == ModeSTARTTLS {
		ldapServer := LdapProtocol + "://" + Hostname + ":" + LdapPort
		return ConnectStartTLS(ldapServer)
		fmt.Println("Connect ModeNonTLS...") // TESTED OK
	}
	return nil
}

func ConnectUnsecureDialURL(ldapServer string) *ldap.Conn {
	l, err := ldap.DialURL(ldapServer)
	if err != nil {
		panic(err)
	}
	//defer l.Close()
	return l
}

func ConnectSecureDialURL(ldapsServer string) *ldap.Conn {
	l, err := ldap.DialURL(ldapsServer, ldap.DialWithTLSConfig(&tls.Config{InsecureSkipVerify: true}))
	if err != nil {
		panic(err)
	}
	//defer l.Close()
	return l
}

func ConnectStartTLS(ldapServer string) *ldap.Conn {
	l, err := ldap.DialURL(ldapServer)
	if err != nil {
		panic(err)
	}
	//defer l.Close()
	err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
	if err != nil {
		panic(err)
	}
	return l
}

func Bind(l *ldap.Conn, user string, pass string) bool {
	fmt.Println("Bind...")
	err := l.Bind(user, pass)
	if err != nil {
		panic(err)
	}
	return true
}

func Close(l *ldap.Conn) bool {
	fmt.Println("Close...")
	l.Close()
	return true
}

func Search(l *ldap.Conn, filter []string, attributes []string) []*ldap.Entry {
	searchRequest := ldap.NewSearchRequest(
		LdapBaseDN,
		ldap.ScopeWholeSubtree, ldap.DerefAlways, 0, 0, false,
		filter[0],
		attributes,
		nil)
	sr, err := l.Search(searchRequest)
	if err != nil {
		panic(err)
	}
	//output := "TestSearch: %s -> num of entries = %d -> " + searchRequest.Filter + " -> " + strconv.Itoa(len(sr.Entries))
	//fmt.Println(output)
	return sr.Entries
}

// ************************************************************

func SearchStartTLS() {

}

func SearcWithPaging() {
	fmt.Println("SearcWithPaging...")
}

func Filter() {
	fmt.Println("Filter...")
}

// NOTE: Admin actions not implemented
func Modify() {
	fmt.Println("Modify...")
}

func Add() {
	fmt.Println("Add...")
}

func Delete() {
	fmt.Println("Delete...")
}

func ModifyDN() {
	fmt.Println("ModifyDN...")
}
