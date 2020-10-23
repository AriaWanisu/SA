import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Button from '@material-ui/core/Button';
import Paper from '@material-ui/core/Paper';
import { DefaultApi } from '../../api/apis/';
import { EntCourseItem} from '../../api/';
import Swal from 'sweetalert2'; // alert
 
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
  const classes = useStyles();
  const api = new DefaultApi();
              
  const [courseitems, setCoursesItem] = useState<EntCourseItem[]>(Array);
  const [loading, setLoading] =  useState(true);

  useEffect(() => {
        const getCoursesItem = async () => {
            const res = await api.listCourseitem({ limit: 20, offset: 0 });
            setLoading(false);
            setCoursesItem(res);
        };
        getCoursesItem();
    }, [loading]);
    
  const Toast = Swal.mixin({
      toast: true,
      position: 'top-end',
      showConfirmButton: false,
      timer: 3000,
      timerProgressBar: true,
      didOpen: toast => {
        toast.addEventListener('mouseenter', Swal.stopTimer);
        toast.addEventListener('mouseleave', Swal.resumeTimer);
      },
  });

  const deleteUsers = async (id: number) => {
    const apiUrl = 'http://localhost:8080/api/v1/CourseItems/';
    const requestOptions = {
      method: 'DELETE',
      headers: { 'Content-Type': 'application/json' },
    };

    fetch(apiUrl + id, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          Toast.fire({
            icon: 'success',
            title: 'ลบสำเร็จ',
          });
        }
      });
      setLoading(true);
  };

  return (
      <TableContainer component={Paper}>
        <Table className={classes.table} aria-label="simple table">
          <TableHead>
            <TableRow>
                <TableCell align="center">No.</TableCell>
                <TableCell align="center">Course name</TableCell>
                <TableCell align="center">Subject id</TableCell>
                <TableCell align="center">Subject Name</TableCell>
                <TableCell align="center">Subject Type</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
              {courseitems.map((item: any) => (
                  <TableRow key={item.id}>
                      <TableCell align="center">{item.id}</TableCell> 
                      <TableCell align="center">{item.edges?.courses?.courseName}</TableCell> 
                      <TableCell align="center">{item.edges?.subjects?.id}</TableCell> 
                      <TableCell align="center">{item.edges?.subjects?.subjectName}</TableCell> 
                      <TableCell align="center">{item.edges?.types?.typeName}</TableCell>
                      <TableCell align="center">
                                <Button
                                      onClick={() => {
                                          deleteUsers(item.id);
                                      }}
                                      style={{ marginLeft: 10 }}
                                      variant="contained"
                                      color="secondary"
                                >
                                    Delete
                                </Button></TableCell> 
                  </TableRow>
              ))}
          </TableBody>
        </Table>
      </TableContainer>
 );
}
