package models

import (
	"echo/databases"
	"net/http"
)

type Pegawai struct {
	Id     int    `json:"id"`
	Nama   string `json:"nama"`
	Alamat string `json:"alamat"`
	Hp     int    `json:"hp"`
}

func FetchAllPegawai() (ResponeMessage, error) {
	var obj Pegawai
	var arrObj []Pegawai
	var res ResponeMessage

	con := databases.CreateCon()
	sql := " select * from data_pegawai"

	rows, err := con.Query(sql)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Nama, &obj.Alamat, &obj.Hp)

		if err != nil {
			return res, err
		}

		arrObj = append(arrObj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj

	return res, nil
}

func StorePegawai(id string, nama string, alamat string, hp string) (ResponeMessage, error) {

	var res ResponeMessage

	con := databases.CreateCon()

	sql := "insert into data_pegawai (id_pegawai, nama, alamat,hp) values (?,?,?,?)"

	stmt, err := con.Prepare(sql)
	if err != nil {
		return res, err
	}

	_, err = stmt.Exec(id, nama, alamat, hp)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]string{
		"lastId": id,
	}
	return res, nil
}

func UpdatePegawai(id int, nama string, alamat string, hp string) (ResponeMessage, error) {

	var res ResponeMessage

	con := databases.CreateCon()

	sql := "update data_pegawai set id_pegawai = ?, nama= ?, alamat=?,hp=? where id_pegawai = ?"

	stmt, err := con.Prepare(sql)
	if err != nil {
		return res, err
	}
	result, err := stmt.Exec(id, nama, alamat, hp, id)
	if err != nil {
		return res, err
	}

	rowaffeted, err := result.RowsAffected()
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rowaffeted": rowaffeted,
	}
	return res, nil

}

func DeletePegawai(id int) (ResponeMessage, error) {
	var res ResponeMessage
	con := databases.CreateCon()

	sql := "delete from data_pegawai where id_pegawai = ?"
	stmt, err := con.Prepare(sql)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)

	if err != nil {
		return res, err
	}

	rowaffeted, err := result.RowsAffected()
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rowaffeted": rowaffeted,
	}
	return res, nil
}
