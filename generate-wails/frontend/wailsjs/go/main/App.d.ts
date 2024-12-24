// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {main} from '../models';

export function DeleteContent(arg1:string,arg2:string):Promise<main.Response>;

export function DeleteTutorial(arg1:string):Promise<main.Response>;

export function GetAllContents(arg1:string):Promise<Array<main.Content>>;

export function GetAllTutorials():Promise<Array<main.Tutorial>>;

export function InsertContent(arg1:main.Content,arg2:string):Promise<main.Response>;

export function InsertTutorial(arg1:main.Tutorial,arg2:Array<number>):Promise<main.Response>;

export function SendGit(arg1:string):Promise<string>;

export function UpdateContent(arg1:main.Content,arg2:string):Promise<main.Response>;

export function UpdateTutorial(arg1:main.Tutorial,arg2:Array<number>):Promise<main.Response>;
