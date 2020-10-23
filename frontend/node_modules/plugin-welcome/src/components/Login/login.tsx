import React, { FC, useEffect} from 'react';
import { Link as RouterLink } from 'react-router-dom';
import Typography from '@material-ui/core/Typography';
import { makeStyles, Theme,createStyles  } from '@material-ui/core/styles';

import {
    Content,
    Header,
    Page,
    pageTheme,
    ContentHeader,
} from '@backstage/core';

import {
    Grid,
    FormControl,
    Select,
    InputLabel,
    MenuItem,
    Button,
    Box,
    TextField,
} from '@material-ui/core';

import Visibility from '@material-ui/icons/Visibility';
import VisibilityOff from '@material-ui/icons/VisibilityOff';
import OutlinedInput from '@material-ui/core/OutlinedInput';
import InputAdornment from '@material-ui/core/InputAdornment';
import IconButton from '@material-ui/core/IconButton';

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
    boxStyle: {
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        flexDirection: 'row'
    },
    formControl: {
        width: 300,
        height: 65,
        marginBottom: theme.spacing(2),
    },
    margin: {
        margin: theme.spacing(1),
      },
      withoutLabel: {
        marginTop: theme.spacing(3),
      },
      textField: {
        width: '25ch',
      },
}));

interface State {
    password: string;
    showPassword: boolean;
  }


const login: FC<{}> = () => {
    const classes = useStyles();
    const [values, setValues] = React.useState<State>({
        password: '',
        showPassword: false,
      });

    const handleChange = (prop: keyof State) => (event: React.ChangeEvent<HTMLInputElement>) => {
        setValues({ ...values, [prop]: event.target.value });
    };
    
      const handleClickShowPassword = () => {
        setValues({ ...values, showPassword: !values.showPassword });
    };
    
      const handleMouseDownPassword = (event: React.MouseEvent<HTMLButtonElement>) => {
        event.preventDefault();
    };

    return (

        <Page theme={pageTheme.home}>
            <Header
                title="ระบบบันทึกสร้างและจัดการหลักสูตร"
                subtitle="สามารถบันทึกและตรวจสอบข้อมูลหลักสูตรได้ที่นี่"
            ></Header>

            <Content>
                <ContentHeader title="เข้าสู่ระบบก่อนใช้งาน">
                    <Box
                        display="flex"
                        justifyContent="flex-start"
                        flexDirection="column"
                    >
                      
                    </Box>

                </ContentHeader>

                <Grid container>
                    <Grid item xs={12}>
                    <Box className={classes.boxStyle}>
                            <Typography variant="subtitle1" gutterBottom style={{ width: 130 }}>
                                User
                            </Typography>
                            <FormControl className={classes.formControl} variant="outlined" >
                                <InputLabel></InputLabel>
                                <TextField
                                    name="USER"
                                    variant="outlined"
                                    autoComplete="off"
                                >
                                </TextField>
                            </FormControl>
                        </Box>
                        <Box className={classes.boxStyle}>
                            <Typography variant="subtitle1" gutterBottom style={{ width: 130 }}>
                                    Password
                            </Typography>
                            <FormControl className={classes.formControl} variant="outlined">
                                <InputLabel htmlFor="outlined-adornment-password">Password</InputLabel>
                                    <OutlinedInput
                                        id="outlined-adornment-password"
                                        type={values.showPassword ? 'text' : 'password'}
                                        value={values.password}
                                        onChange={handleChange('password')}
                                        endAdornment={
                                            <InputAdornment position="end">
                                                <IconButton
                                                    aria-label="toggle password visibility"
                                                    onClick={handleClickShowPassword}
                                                    onMouseDown={handleMouseDownPassword}
                                                    edge="end"
                                                    >
                                                    {values.showPassword ? <Visibility /> : <VisibilityOff />}
                                                </IconButton>
                                            </InputAdornment>
                                        }
                                        labelWidth={70}
                                    />
                            </FormControl>
                        </Box>
                        <Box
                            display="flex"
                            justifyContent="center" >
                            <Button
                                style={{ width: 300, height: 40, marginLeft: 130 }}
                                variant="contained"
                                color="primary"
                                component={RouterLink}
                                to="/WelcomePage"
                            >
                                เข้าสู่ระบบ
                            </Button>

                        </Box>
                    </Grid>
                </Grid>
            </Content>
        </Page>
    );
};

export default login;