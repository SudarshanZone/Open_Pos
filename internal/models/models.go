package models

//st_usr_prf1
type St_Usr_Prfl struct {
	C_User_Id          string    // c_user_id[9]
	L_Session_Id       int64     // l_session_id
	C_Pipe_Id          string    // c_pipe_id[3]
	C_Cln_Mtch_Accnt   string    // c_cln_mtch_accnt[11]
	C_C_DpId           int64     // C_c_dp_id
	C_C_Dp_ID          [5]string // c_dp_id[5][9]
	C_Dp_Clnt_Id       [5]string // c_dp_clnt_id[5][9]
	C_C_Bnk_Accnt_Nmbr int64     // C_c_bnk_accnt_nmbr
	C_Bnk_Accnt_Nmbr   [5]string // c_bnk_accnt_nmbr[5][21]
	L_Clnt_Ctgry       int64     // l_clnt_ctgry
	L_Usr_Flg          int64     // l_usr_flg
	C_Rout_Crt         string    // c_rout_crt[4]
}


// sterrmsg
type St_Err_Msg struct {
	C_Err_No   string //maxlength 7
	C_Err_Msg  string // 256 length max
	C_Rout_Crt string //maxlength 4
}





