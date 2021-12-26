import { Get, HttpStatus, Injectable } from '@nestjs/common';
import { Client } from '@notionhq/client';
import { DataDto } from './dto/data.dto';

@Injectable()
export class DataService {

    async findNotion(body: DataDto): Promise<any> {
        const { service, startdate, enddate } = body;
        const notion = new Client({
            auth: process.env.NOTION_SECRET
        });
        const res = notion.databases.query({
            database_id: process.env.NOTION_DB,
            filter: {
                property: "FindDate",
                date: {
                    after: startdate,
                    before: enddate
                }
            }
        });
        const dbdata = (await res).results.map((page:any) => {
            const id = page.id;
            // const des = notion.databases.retrieve({
            //     page_id : id
            // })
            return {
                type : page.properties.Type.select.name,
                finddate : page.properties.FindDate.date.start,
                vulnerability : page.properties.Vulnerability.multi_select[0].name,
                target : page.properties.Target.select.name,
                human : page.properties.Human.people[0].name,
                title : page.properties.Title.title[0].plain_text,
            }
        })
        //console.log(dbdata);
        return dbdata;
    }
}
